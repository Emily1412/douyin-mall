package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ServiceHealth = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_health_status",
			Help: "Health status of services (1 for healthy, 0 for unhealthy)",
		},
		[]string{"service"},
	)
)

// UpdateHealthMetrics 更新服务健康状态指标
func UpdateHealthMetrics(services map[string]bool) {
	for service, healthy := range services {
		value := 0.0
		if healthy {
			value = 1.0
		}
		ServiceHealth.WithLabelValues(service).Set(value)
	}
}
