package database

type Repository interface {
	SaveCPUMetrics() error
	SaveMemMetrics() error
}
