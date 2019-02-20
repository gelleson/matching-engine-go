package main

import (
	"fmt"
	"gitlab.com/gelleson/trader/matching-engine-go/src/core"
	"time"
)

func main() {
	engine := core.Engine{}
	cron := core.Cron{
		Mask:[]core.Day{
			{
				Days: time.Wednesday,
				Hour: 13,
				Minute: 39,
				Second: 10,
			},
		},
		Repeat: false,
		Serve: func() {
			fmt.Println("admin")
		},
	}
	engine.SetCron(cron)
	engine.Serve()
}