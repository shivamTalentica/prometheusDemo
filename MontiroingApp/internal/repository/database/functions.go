package database

import (
	"time"
)

func (DB *DBRepository) SaveCPUMetrics(host string, appName string, cpuUse float32) error {
	query := `INSERT INTO CPUUsage  (host, appname, cpu_usage, created_at) values ($1,$2,$3)`

	DB.db.QueryRow(query, host, appName, cpuUse, time.Now())
	return nil
}

func (DB *DBRepository) SaveMemMetrics(host string, appName string, memUse float32) error {
	query := `INSERT INTO MemUsage (host, appname, memory_usage, created_at) VALUES ($1, $2, $3, $4))`

	DB.db.QueryRow(query, host, appName, memUse, time.Now())
	return nil
}
