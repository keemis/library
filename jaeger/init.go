package jaeger

import (
	"github.com/astaxie/beego"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var (
	Tracer Trace
)

type Trace struct {
	used   bool
	tracer opentracing.Tracer
}

// Init 初始化
func Init(conf *jaegercfg.Configuration) {
	if conf == nil {
		conf = &jaegercfg.Configuration{
			ServiceName: beego.AppConfig.DefaultString("JaegerServiceName", "UnknownApp"),
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: beego.AppConfig.DefaultString("JaegerLocalAgent", "127.0.0.1:6831"),
			},
		}
	}
	trace, _, err := (*conf).NewTracer()
	if err != nil {
		panic(err)
	}
	Tracer.tracer = trace
	Tracer.used = true
}
