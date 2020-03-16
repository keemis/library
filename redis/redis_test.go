package redis

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// init
	Init(&Option{
		Sentinel:     false,
		Addrs:        []string{"127.0.0.1:6379"},
		Password:     "",
		DB:           0,
		DialTimeout:  time.Duration(3) * time.Second,
		ReadTimeout:  time.Duration(2) * time.Second,
		WriteTimeout: time.Duration(2) * time.Second,
		PoolSize:     500,
		PoolTimeout:  time.Duration(3) * time.Second,
		MasterName:   "master",
	})

	// set
	Client().Set("k", "v", 0)
	v, err := Client().Get("k").Result()
	t.Logf("value: %v, err: %v", v, err)

	// del
	affected, err := Client().Del("k").Result()
	t.Logf("affected: %v, err: %v", affected, err)

	// get
	v, err = Client().Get("k").Result()
	t.Logf("value: %v, err: %v", v, err)
}
