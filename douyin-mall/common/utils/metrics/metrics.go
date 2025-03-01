// 指标
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	

	// RPC 请求计数器
	RPCRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rpc_requests_total",
			Help: "Total number of RPC requests",
		},
		[]string{"method", "status"},
	)

	// RPC 处理时间
	RPCDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "rpc_duration_seconds",
			Help:    "RPC processing duration in seconds",
			Buckets: prometheus.LinearBuckets(0, 0.1, 10),
		},
		[]string{"method"},
	)
)
