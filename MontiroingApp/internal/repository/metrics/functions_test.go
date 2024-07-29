package metrics

import (
	"reflect"
	"testing"
)

func TestMetricsRepository_GetCPUUSages(t *testing.T) {

	type fields struct {
		hosturl  string
		memQuery string
		cpuQuery string
	}
	var pr PrometheusResponse

	tests := []struct {
		name    string
		fields  fields
		want    PrometheusResponse
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				hosturl:  "123.45.67:9000",
				memQuery: "mem",
				cpuQuery: "cpu",
			},
			want:    pr,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetricsRepository{
				hosturl:  tt.fields.hosturl,
				memQuery: tt.fields.memQuery,
				cpuQuery: tt.fields.cpuQuery,
			}
			got, err := m.GetCPUUSages()
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsRepository.GetCPUUSages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetricsRepository.GetCPUUSages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetricsRepository_GetMemoryUsages(t *testing.T) {
	type fields struct {
		hosturl  string
		memQuery string
		cpuQuery string
	}
	var pr PrometheusResponse

	tests := []struct {
		name    string
		fields  fields
		want    PrometheusResponse
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				hosturl:  "123.45.67:9000",
				memQuery: "mem",
				cpuQuery: "cpu",
			},
			want:    pr,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetricsRepository{
				hosturl:  tt.fields.hosturl,
				memQuery: tt.fields.memQuery,
				cpuQuery: tt.fields.cpuQuery,
			}
			got, err := m.GetMemoryUsages()
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsRepository.GetMemoryUsages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetricsRepository.GetMemoryUsages() = %v, want %v", got, tt.want)
			}
		})
	}
}
