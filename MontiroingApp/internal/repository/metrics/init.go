package metrics

func InitPromRetriver() MetricsRepository {
	// Todo: Add host in config file
	var m MetricsRepository
	m.hosturl = "http://localhost:9090/api/v1/query?query="
	m.cpuQuery = "app_cpu_usage_percent"
	m.memQuery = "app_memory_usage_bytes"

	return m
}
