package usecase

import (
	"testing"

	"github.com/shivam/promotheus/internal/repository/database"
	"github.com/shivam/promotheus/internal/repository/metrics"
)

func TestUsecase_FetchAndSaveMetrics(t *testing.T) {
	type fields struct {
		DB      database.DBRepository
		Metrics metrics.MetricsRepository
	}
	tests := []struct {
		name   string
		fields fields
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				DB:      tt.fields.DB,
				Metrics: tt.fields.Metrics,
			}
			u.FetchAndSaveMetrics()
		})
	}
}
