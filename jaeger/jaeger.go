package jaeger

import (
	"errors"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// IsUsed 是否启用jaeger
func (t Trace) IsUsed() bool {
	return t.used
}

// StartSpan 上传到jaeger
func (t Trace) StartSpan(head http.Header, operationName string, baggage map[string]string) error {
	if !t.IsUsed() {
		return errors.New("please initialize first")
	}
	carrier := opentracing.HTTPHeadersCarrier(head)
	sm, err := t.tracer.Extract(opentracing.HTTPHeaders, carrier)
	if err != nil {
		return err
	}
	span := t.tracer.StartSpan(operationName, ext.RPCServerOption(sm))
	for k, v := range baggage {
		span = span.SetBaggageItem(k, v)
	}
	span.Finish()
	return nil
}
