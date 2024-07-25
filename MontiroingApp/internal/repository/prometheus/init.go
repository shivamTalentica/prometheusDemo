package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func InitPrometheusMetrics() PromRepository {
	var p PromRepository
	p.cpuUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_cpu_usage_percent",
		Help: "Current CPU usage percentage",
	})
	p.memUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_usage_bytes",
		Help: "Current memory usage in bytes",
	})
	p.Port = ":8081"
	return p
}
