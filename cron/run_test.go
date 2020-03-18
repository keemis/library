package cron

import (
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

//
// https://github.com/robfig/cron
//

func TestRun(t *testing.T) {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		t.Logf("Every Minute ")
	})

	c.Start()

	time.Sleep(1 * time.Hour)
}
