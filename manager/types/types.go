package types

import (
	"context"
	"math/big"

	"github.com/WJX2001/vrf-node-new/sign"
)

type SignMsgRequest struct {
	BlockNumber *big.Int `json:"block_number"`
	TxHash      []byte   `json:"tx_hash"`
	TxType      string   `json:"tx_type"`
	StateRoot   string   `json:"state_root"`
}

type NodeSignRequest struct {
	Timestamp   int64       `json:"timestamp"`
	Nodes       []string    `json:"nodes"`
	RequestBody interface{} `json:"request_body"`
}

type SignMsgResponse struct {
	Signature       []byte `json:"signature"`
	G2Point         []byte `json:"g2_point"`
	NonSignerPubkey []byte `json:"non_signer_pubkey"`
	Vote            uint8  `json:"vote"`
}

type SignResult struct {
	Signature        *sign.G1Point   `json:"signature"`
	G2Point          *sign.G2Point   `json:"g2_point"`
	NonSignerPubkeys []*sign.G1Point `json:"non_signer_pubkeys"`
}

type Method string

const (
	SignMsgBatch Method = "signMsgBatch"
)

func (m Method) String() string {
	return string(m)
}

// Context ---------------------------------------------
type Context struct {
	ctx            context.Context
	requestId      string
	availableNodes []string
	approvers      []string
	unApprovers    []string
	electionId     uint64
	stateBatchRoot string
}

func NewContext() Context {
	return Context{
		ctx: context.Background(),
	}
}

func (c Context) RequestId() string {
	return c.requestId
}

func (c Context) AvailableNodes() []string {
	return c.availableNodes
}

func (c Context) Approvers() []string {
	return c.approvers
}

func (c Context) WithRequestId(requestId string) Context {
	c.requestId = requestId
	return c
}

func (c Context) WithAvailableNodes(nodes []string) Context {
	c.availableNodes = nodes
	return c
}
