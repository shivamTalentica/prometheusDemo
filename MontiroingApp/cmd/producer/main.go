package main

import (
	"fmt"

	"github.com/shivam/promotheus/internal/repository/prometheus"
)

func main() {

	prom := prometheus.InitPrometheusMetrics()

	fmt.Println(prom)
	prom.RecordMetrics()
	prom.StartServer()
}
