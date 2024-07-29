package metrics

import (
	"reflect"
	"testing"
)

func TestInitPromRetriver(t *testing.T) {

	tests := []struct {
		name string
		want MetricsRepository
	}{
		{
			name: "Test Case 1",
			want: MetricsRepository{
				hosturl:  "http://localhost:9090/api/v1/query?query=",
				memQuery: "app_memory_usage_bytes",
				cpuQuery: "app_cpu_usage_percent",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitPromRetriver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitPromRetriver() = %v, want %v", got, tt.want)
			}
		})
	}
}
