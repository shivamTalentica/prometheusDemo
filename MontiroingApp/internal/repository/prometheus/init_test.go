package prometheus

import (
	"reflect"
	"testing"
)

func TestInitPrometheusMetrics(t *testing.T) {

	tests := []struct {
		name string
		want PromRepository
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitPrometheusMetrics(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitPrometheusMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}
