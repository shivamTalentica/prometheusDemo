package prometheus

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func (P PromRepository) StartServer() error {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(P.Port, nil)
	return nil
}

func (P PromRepository) RecordMetrics() error {

	go func() {
		for {
			err := P.LogCPUMetrics()
			if err != nil {

			}
			err = P.LogMemMetrcics()
			if err != nil {

			}
			time.Sleep(4 * time.Second)
		}
	}()
	return nil
}

func (P PromRepository) LogCPUMetrics() error {
	percent, _ := cpu.Percent(0, false)
	if len(percent) > 0 {
		P.cpuUsage.Set(percent[0])
	}
	fmt.Println("cpu used", percent[0])
	return nil
}

func (P PromRepository) LogMemMetrcics() error {
	v, _ := mem.VirtualMemory()
	P.memUsage.Set(float64(v.Used))
	fmt.Println("Memory used", v.Used)
	return nil
}
