package usecase

import (
	"github.com/shivam/promotheus/internal/repository/database"
	"github.com/shivam/promotheus/internal/repository/metrics"
)

type Usecase struct {
	DB      database.DBRepository
	Metrics metrics.MetricsRepository
}
