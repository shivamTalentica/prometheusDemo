package prometheus

import "github.com/prometheus/client_golang/prometheus"

type PromRepository struct {
	cpuUsage prometheus.Gauge
	memUsage prometheus.Gauge
	Port     string
}
