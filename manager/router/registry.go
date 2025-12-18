package router

import (
	"errors"
	"net/http"

	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/manager/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Registry struct {
	signService types.SignService
	db          *database.DB
}

func NewRegistry(signService types.SignService, db *database.DB) *Registry {
	return &Registry{
		signService: signService,
		db:          db,
	}
}

func (registry *Registry) SignMsgHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.SignMsgRequest
		// 解析JSON 请求体
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}

		// 验证必填字段
		if len(request.TxHash) == 0 || request.BlockNumber == nil || request.TxType == "" {
			c.JSON(http.StatusBadRequest, errors.New("tx_hash, block_number and tx_type must not be nil"))
			return
		}

		// 调用签名服务
		var result *types.SignResult
		var err error
		result, err = registry.signService.SignMsgBatch(request)

		if err != nil {
			c.String(http.StatusInternalServerError, "failed to sign msg")
			log.Error("failed to sign msg", "error", err)
			return
		}
		// 写入响应体
		if _, err = c.Writer.Write(result.Signature.Serialize()); err != nil {
			log.Error("failed to write signature to response writer", "error", err)
		}
	}
}

func (registry *Registry) PrometheusHandler() gin.HandlerFunc {
	// 创建 Prometheus HTTP 处理器
	h := promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(
			prometheus.DefaultGatherer,                   // 指标收集器
			promhttp.HandlerOpts{MaxRequestsInFlight: 3}, // 最多同时处理3个请求，用于限流
		),
	)

	// 返回 Gin 处理器函数
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
