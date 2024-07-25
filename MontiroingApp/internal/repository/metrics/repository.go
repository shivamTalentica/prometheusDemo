package metrics

type Repository interface {
	GetCPUUSages() (PrometheusResponse, error)
	GetMemoryUsages() (PrometheusResponse, error)
}
