package usecase

import (
	"fmt"
	"time"

	"github.com/shivam/promotheus/internal/utils"
)

func (u Usecase) FetchAndSaveMetrics() {

	go func() {
		for {
			cpuResult, err := u.Metrics.GetCPUUSages()
			if err != nil {
				return
			}
			//fmt.Println("cp")
			if len(cpuResult.Data.Result) > 0 && len(cpuResult.Data.Result[0].Value) > 1 {

				cpuUse := cpuResult.Data.Result[0].Value[1].(string)
				instance := cpuResult.Data.Result[0].Metric["instance"]
				appName := cpuResult.Data.Result[0].Metric["job"]
				cpuFloat := utils.StringTofloat(cpuUse)

				fmt.Println("cpu usages", cpuFloat, " ", instance)

				u.DB.SaveCPUMetrics(instance, appName, cpuFloat)
			}

			// Todo: comeup with better way, like ticker or something
			time.Sleep(8 * time.Second)
			//	break
		}
	}()

	go func() {
		for {
			memResult, err := u.Metrics.GetMemoryUsages()
			if err != nil {
				return
			}

			if len(memResult.Data.Result) > 0 && len(memResult.Data.Result[0].Value) > 0 {

				memUse := memResult.Data.Result[0].Value[1].(string)
				instance := memResult.Data.Result[0].Metric["instance"]
				//	appName := memResult.Data.Result[0].Metric["job"]
				memFloat := utils.StringTofloat(memUse)

				fmt.Println("Memory usages", memFloat, " ", instance)
			}

			// Todo: comeup with better way, like ticker or something
			time.Sleep(8 * time.Second)
			break
		}
	}()
}
