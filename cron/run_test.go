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
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* * * * * *", func() {
		t.Logf("Every Second ")
	})

	c.Start()

	time.Sleep(1 * time.Hour)
}
