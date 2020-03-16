package limiter

import (
	"errors"
	"time"

	"golang.org/x/time/rate"
)

var (
	limList = map[string]*Lim{} // 限频规则列表
)

// Register 注册限流业务和规则
// eg:
//    limiters.Register("ip_limit", limiters.New(1000*time.Millisecond, 5))
func Register(limName string, lim *Lim) error {
	if limName == "" || lim == nil {
		return errors.New("params error")
	}
	limList[limName] = lim
	return nil
}

// GetLimiter 获取Limiter
// eg:
//    limiter := limiters.GetLimiter("ip_limit", "192.168.0.1")
func GetLimiter(limName string, key string) *rate.Limiter {
	if lim, ok := limList[limName]; ok {
		return lim.GetLimiter(key)
	}
	return nil
}

// AllowN 是否允许执行
// eg:
//    allow := limiters.AllowN("ip_limit", "192.168.0.1")
func AllowN(limName string, key string, n ...int) bool {
	cnt := 1
	if len(n) == 1 && n[0] != 0 {
		cnt = n[0]
	}
	limiter := GetLimiter(limName, key)
	if limiter == nil {
		return false
	}
	return limiter.AllowN(time.Now(), cnt)
}
