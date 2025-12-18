package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/WJX2001/vrf-node-new/manager/types"
	"github.com/WJX2001/vrf-node-new/sign"
	"github.com/WJX2001/vrf-node-new/ws/server"

	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (m *Manager) sign(ctx types.Context, request interface{}, method types.Method) (types.SignResult, error) {
	// 创建响应通道 respChan 和 停止通道 stopChan
	// 在 websocket 服务器注册响应通道，
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})

	if err := m.wsServer.RegisterResChannel(ctx.RequestId(), respChan, stopChan); err != nil {
		m.log.Error("failed to register response channel at signing step", "err", err)
		return types.SignResult{}, err
	}

	m.log.Info("Registered ResChannel with requestID", "requestID", ctx.RequestId())

	errSendChan := make(chan struct{})
	responseNodes := make(map[string]struct{})
	var err error
	var respNumber int
	var validSignResult types.SignResult
	var g2Point *sign.G2Point
	var g2Points []*sign.G2Point // 收集的 G2 点
	var g1Point *sign.G1Point
	var g1Points []*sign.G1Point         // 收集的 G2 点
	var NonSignerPubkeys []*sign.G1Point // 未签名节点的公钥

	// 启动响应收集协程
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cctx, cancel := context.WithTimeout(context.Background(), m.signTimeout)
		defer func() {
			m.log.Info("exit signing process")
			cancel()
			close(stopChan)
			wg.Done()
		}()

		for {
			select {
			case <-errSendChan:
				return
			case resp := <-respChan:
				m.log.Info(fmt.Sprintf("signed response: %s", resp.RpcResponse.String()), "node", resp.SourceNode)
				if !ExistsIgnoreCase(ctx.AvailableNodes(), resp.SourceNode) {
					// 忽略该消息，发送者不应被纳入审批者集合
					continue
				}
				respNumber++
				func() {
					// 使用 defer 确保记录响应节点
					// 无论函数如何返回，都会将 resp.SourceNode 记录到 responseNodes map 中，用于追踪已响应的节点
					defer func() {
						responseNodes[resp.SourceNode] = struct{}{}
					}()

					/*
						处理签名响应的三种情况：
						 1：错误响应：记录错误信息
						 2：Vote != 1 未签名：将 NonSignerPubkey 添加到 NonSignerPubkeys
						 3: Vote == 1 已签名：解析 G2Point 和 signature 分别添加到 g2Points 和 g1Points
					*/
					if resp.RpcResponse.Error != nil {
						m.log.Error("Unrecognized error code",
							"err_code", resp.RpcResponse.Error.Code,
							"err_data", resp.RpcResponse.Error.Data,
							"err_message", resp.RpcResponse.Error.Message)
						return
					} else {
						var signResponse types.SignMsgResponse
						if err = tmjson.Unmarshal(resp.RpcResponse.Result, &signResponse); err != nil {
							m.log.Error("failed to unmarshal sign response", "err", err)
							return
						}

						if signResponse.Vote != 1 { // 0: 没有签名 1:代表签名
							// G1: 椭圆曲线 G1 群的点，用于BLS 签名
							// NonSignerPubkey: 未签名节点的公钥
							// Deserialize: 将序列化的字节数据还原为可用的 G1Point 对象
							g1Point, err = new(sign.G1Point).Deserialize(signResponse.NonSignerPubkey)
							if err != nil {
								m.log.Error("failed to deserialize g1Point", "err", err)
								return
							}
							NonSignerPubkeys = append(NonSignerPubkeys, g1Point)
							return
						}

						dG2Point, err := g2Point.Deserialize(signResponse.G2Point)
						if err != nil {
							m.log.Error("failed to deserialize g2Point", "err", err)
							return
						}

						dSign, err := g1Point.Deserialize(signResponse.Signature)
						if err != nil {
							m.log.Error("failed to deserialize signature", "err", err)
							return
						}

						g2Points = append(g2Points, dG2Point)
						g1Points = append(g1Points, dSign)
						return
					}
				}()

			case <-cctx.Done():
				m.log.Warn("wait for signature timeout", "requestId", ctx.RequestId(), "received responses len", respNumber)
				return
			default:
				if respNumber == len(ctx.AvailableNodes()) {
					m.log.Info("received all signing responses", "requestId", ctx.RequestId(), "received responses len", respNumber)
					return
				}
			}
		}
	}()

	m.sendToNodes(ctx, request, method, errSendChan)
	wg.Wait()

	if respNumber < len(ctx.AvailableNodes())*2/3 {
		return validSignResult, errNotEnoughVoteNode
	}
	aSign, aG2Point := aggregateSignaturesAndG2Point(g1Points, g2Points)
	if aSign != nil {
		validSignResult = types.SignResult{
			NonSignerPubkeys: NonSignerPubkeys,
			Signature:        aSign,
			G2Point:          aG2Point,
		}
	}

	return validSignResult, nil
}

func (m *Manager) sendToNodes(ctx types.Context, request interface{}, method types.Method, errSendChan chan struct{}) {
	nodes := ctx.AvailableNodes()
	nodeRequest := types.NodeSignRequest{
		Timestamp:   time.Now().UnixMilli(),
		Nodes:       ctx.Approvers(),
		RequestBody: request,
	}

	requestBz, err := json.Marshal(nodeRequest)
	if err != nil {
		m.log.Error("failed to json marshal node request", "err", err)
		errSendChan <- struct{}{}
		return
	}
	rpcRequest := tmtypes.NewRPCRequest(tmtypes.JSONRPCStringID(ctx.RequestId()), method.String(), requestBz)
	for _, node := range nodes {
		go func(node string, request tmtypes.RPCRequest) {
			if err := m.wsServer.SendMsg(
				server.RequestMsg{
					RpcRequest: request,
					TargetNode: node,
				}); err != nil {
				m.log.Error("failed to send sign request to nodes", "err", err)
				errSendChan <- struct{}{}
				return
			}
		}(node, rpcRequest)
	}
}

// 聚合签名
func aggregateSignaturesAndG2Point(signatures []*sign.G1Point, points []*sign.G2Point) (*sign.G1Point, *sign.G2Point) {
	if len(signatures) == 0 {
		return nil, nil
	}

	var aggSig *sign.G1Point
	var g2Point *sign.G2Point

	for _, sig := range signatures {
		if aggSig == nil {
			aggSig = sig.Clone()
		} else {
			aggSig.Add(sig)
		}
	}

	for _, point := range points {
		if g2Point == nil {
			g2Point = point.Clone()
		} else {
			g2Point.Add(point)
		}
	}
	return aggSig, g2Point
}
