package curl

import (
	"github.com/astaxie/beego/httplib"
	"github.com/keemis/library/logs"
)

// Curl 结构体
type Curl struct {
	*httplib.BeegoHTTPRequest
	log       logs.Logger
	used      bool
	tlsSecure bool
}

// Option 配置选项
type Option func(*Curl)

// OptLogger 设置日志对象
func OptLogger(log logs.Logger) Option {
	return func(curl *Curl) {
		curl.log = log
	}
}

// OptTlsSecure 设置是否验证TLS证书
func OptTlsSecure(tlsSecure bool) Option {
	return func(curl *Curl) {
		curl.tlsSecure = tlsSecure
	}
}

// New 创建对象
func New(options ...Option) Curl {
	curl := Curl{
		tlsSecure: true,
	}
	for _, option := range options {
		option(&curl)
	}
	return curl
}

// 请求类型
type requestType = int

// 请求类型枚举
const (
	typeGet  requestType = 1
	typePost requestType = 2
)
