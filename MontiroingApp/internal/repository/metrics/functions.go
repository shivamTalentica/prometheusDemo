package metrics

import (
	"encoding/json"
	"log"
	"net/http"
)

func (m MetricsRepository) GetCPUUSages() (PrometheusResponse, error) {

	reqUrl := m.hosturl + m.cpuQuery
	var pr PrometheusResponse

	resp, err := http.Get(reqUrl)
	if err != nil {
		return pr, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from Prometheus: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	if pr.Status != "success" {
		log.Fatalf("Unsuccessful response from Prometheus: %s", pr.Status)
	}

	return pr, nil
}

func (m MetricsRepository) GetMemoryUsages() (PrometheusResponse, error) {

	reqUrl := m.hosturl + m.memQuery
	var pr PrometheusResponse

	resp, err := http.Get(reqUrl)
	if err != nil {
		return pr, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from Prometheus: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	if pr.Status != "success" {
		log.Fatalf("Unsuccessful response from Prometheus: %s", pr.Status)
	}

	return pr, nil
}
