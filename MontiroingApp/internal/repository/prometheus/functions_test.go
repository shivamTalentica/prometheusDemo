package prometheus

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestPromRepository_RecordMetrics(t *testing.T) {
	type fields struct {
		cpuUsage prometheus.Gauge
		memUsage prometheus.Gauge
		Port     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			P := PromRepository{
				cpuUsage: tt.fields.cpuUsage,
				memUsage: tt.fields.memUsage,
				Port:     tt.fields.Port,
			}
			if err := P.RecordMetrics(); (err != nil) != tt.wantErr {
				t.Errorf("PromRepository.RecordMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPromRepository_LogCPUMetrics(t *testing.T) {
	type fields struct {
		cpuUsage prometheus.Gauge
		memUsage prometheus.Gauge
		Port     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			P := PromRepository{
				cpuUsage: tt.fields.cpuUsage,
				memUsage: tt.fields.memUsage,
				Port:     tt.fields.Port,
			}
			if err := P.LogCPUMetrics(); (err != nil) != tt.wantErr {
				t.Errorf("PromRepository.LogCPUMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPromRepository_LogMemMetrcics(t *testing.T) {
	type fields struct {
		cpuUsage prometheus.Gauge
		memUsage prometheus.Gauge
		Port     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			P := PromRepository{
				cpuUsage: tt.fields.cpuUsage,
				memUsage: tt.fields.memUsage,
				Port:     tt.fields.Port,
			}
			if err := P.LogMemMetrcics(); (err != nil) != tt.wantErr {
				t.Errorf("PromRepository.LogMemMetrcics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
