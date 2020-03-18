package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/keemis/library/logs"
	"github.com/robfig/cron/v3"
)

//
// https://github.com/robfig/cron
//

type SyncJob struct {
	log logs.Logger
}

func (s SyncJob) Run() {
	fmt.Println("doing some sync ...")
	time.Sleep(5 * time.Second)
}

type Log struct {
}

func (l Log) Info(msg string, keysAndValues ...interface{}) {
	fmt.Println(msg, keysAndValues)
}

func (l Log) Error(err error, msg string, keysAndValues ...interface{}) {
	fmt.Println(err, msg, keysAndValues)
}

func TestRun(t *testing.T) {
	var l Log
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(l),            // 异常拉起
			cron.SkipIfStillRunning(l), // 任务还在跑，就放弃本次
		),
	)

	syncJob := SyncJob{
		log: logs.New(),
	}
	c.AddJob("* * * * * *", syncJob)

	c.AddFunc("* * * * * *", func() {
		fmt.Println("Every Second")
	})

	c.Start()

	time.Sleep(time.Hour)
}
