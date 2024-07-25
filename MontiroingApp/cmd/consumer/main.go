package main

import (
	"log"
	"sync"

	"github.com/shivam/promotheus/internal/repository/database"
	"github.com/shivam/promotheus/internal/repository/metrics"
	"github.com/shivam/promotheus/internal/usecase"
)

func main() {

	var wg sync.WaitGroup
	promMetricsRetrive := metrics.InitPromRetriver()

	dbObj, err := database.InitDB()
	if err != nil {
		log.Fatal("Unable to connect withDB")
	}

	var usecase usecase.Usecase

	usecase.DB = *dbObj
	usecase.Metrics = promMetricsRetrive

	wg.Add(2)
	usecase.FetchAndSaveMetrics()

	wg.Wait()
}
