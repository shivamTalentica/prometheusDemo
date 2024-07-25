package prometheus

type Repository interface {
	StartServer() error
	LogCPUMetrics() error
	LogMemMetrcics() error
	RecordMetrics() error
}
