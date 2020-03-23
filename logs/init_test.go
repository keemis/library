package logs

import (
	"testing"
)

func TestUse(t *testing.T) {
	// 初始化日志组件
	Init(nil)
	// 生成一个新的TraceID
	log := New()
	log.Info("have a bug: %v", "php")
	log.Error("have a bug: %v", "golang")
}

func TestWithTraceID(t *testing.T) {
	// 初始化日志组件
	Init(nil)
	// 自定义TraceID
	log := New(WithTraceID("158493062888748290"))
	log.Info("have a bug: %v", "php")
	log.Error("have a bug: %v", "golang")
}
