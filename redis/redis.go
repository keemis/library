package redis

import (
	"time"

	"github.com/astaxie/beego"
	gRedis "github.com/go-redis/redis"
)

var (
	client *gRedis.Client
)

// Client get redis client
func Client() *gRedis.Client {
	return client
}

// Option config
type Option struct {
	Sentinel     bool
	Addrs        []string
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
	MasterName   string
}

// Init initial redis
// if opt is nil, use beego ini config
func Init(opt *Option) {
	if opt == nil {
		opt = &Option{
			Sentinel:     beego.AppConfig.DefaultBool("RedisSentinel", true),
			Addrs:        beego.AppConfig.DefaultStrings("RedisAddrs", []string{"127.0.0.1:6379", "127.0.0.1:6380"}),
			Password:     beego.AppConfig.DefaultString("RedisPassword", "123456"),
			DB:           beego.AppConfig.DefaultInt("RedisDB", 0),
			DialTimeout:  time.Duration(beego.AppConfig.DefaultInt("RedisDialTimeout", 3)) * time.Second,
			ReadTimeout:  time.Duration(beego.AppConfig.DefaultInt("RedisReadTimeout", 2)) * time.Second,
			WriteTimeout: time.Duration(beego.AppConfig.DefaultInt("RedisWriteTimeout", 2)) * time.Second,
			PoolSize:     beego.AppConfig.DefaultInt("RedisPoolSize", 500),
			PoolTimeout:  time.Duration(beego.AppConfig.DefaultInt("RedisPoolTimeout", 3)) * time.Second,
			MasterName:   beego.AppConfig.DefaultString("RedisMasterName", "master"),
		}
	}
	if len(opt.Addrs) == 0 {
		panic("redis config Addrs is empty")
	}
	if opt.Sentinel {
		initSentinel(opt)
	} else {
		initRedis(opt)
	}
}

// initRedis initial redis
func initRedis(opt *Option) {
	c := gRedis.NewClient(&gRedis.Options{
		Addr:         opt.Addrs[0],
		Password:     opt.Password,
		DB:           opt.DB,
		MaxRetries:   2,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		PoolTimeout:  opt.PoolTimeout,
	})
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
	client = c
}

// initSentinel initial redis sentinel
func initSentinel(opt *Option) {
	c := gRedis.NewFailoverClient(&gRedis.FailoverOptions{
		MasterName:    opt.MasterName,
		SentinelAddrs: opt.Addrs,
		OnConnect: func(cn *gRedis.Conn) error {
			return nil
		},
		Password:     opt.Password,
		DB:           opt.DB,
		MaxRetries:   2,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		PoolTimeout:  opt.PoolTimeout,
	})
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
	client = c
}
