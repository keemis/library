package limiter

import (
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

// New is a constructor for Limiters
// interval : 多久产生一个令牌，如：1000 * time.Millisecond
// burst : 令牌桶大小，默认无特殊填 1，建议值 5
func New(interval time.Duration, burst int) *Lim {
	if burst <= 0 {
		burst = 1
	}
	lmt := &Lim{}
	lmt.interval = interval
	lmt.burst = burst
	lmt.tokenBuckets = gocache.New(30*time.Second, 30*time.Second)
	return lmt
}

// Lim is a config struct to limit a particular request handler.
type Lim struct {
	interval     time.Duration
	burst        int
	tokenBuckets *gocache.Cache
	sync.RWMutex
}

// GetLimiter 获取一个Limiter
func (l *Lim) GetLimiter(key string) *rate.Limiter {
	l.Lock()
	defer l.Unlock()
	limiterInf, found := l.tokenBuckets.Get(key)
	if !found {
		limiterInf = rate.NewLimiter(rate.Every(l.interval), l.burst)
	}
	bucketTTL := l.interval + 10*time.Second
	l.tokenBuckets.Set(key, limiterInf, bucketTTL)
	return limiterInf.(*rate.Limiter)
}
