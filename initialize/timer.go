package initialize

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"yuyu/config"
	"yuyu/global"
	"yuyu/utils"
)

func Timer() {
	if global.GvaConfig.Timer.Start {
		for i := range global.GvaConfig.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.GvaConfig.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.GvaTimer.AddTaskByFunc("ClearDB", global.GvaConfig.Timer.Spec, func() {
					err := utils.ClearTable(global.GvaDb, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.GvaConfig.Timer.Detail[i])
		}
	}
}
