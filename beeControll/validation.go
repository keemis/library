package beeControll

import (
	"encoding/json"
	"reflect"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/keemis/library/validation"
)

// ValidForm 校验、解析Query、Form、Body到Po
func (u *BaseController) ValidForm(po interface{}) {
	rv := reflect.ValueOf(po)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		u.ApiError("param must be ptr", -1000)
	}
	// encode Query、Form、Body
	byts, err := json.Marshal(u.bodyStore)
	if err != nil {
		u.ApiErrorf("failed to encode form: %v", err)
	}
	// use jsoniter ext
	if len(byts) > 2 {
		extra.RegisterFuzzyDecoders()
		extra.RegisterTimeAsInt64Codec(time.Second)
		if err := jsoniter.Unmarshal(byts, po); err != nil {
			u.ApiErrorf("failed to decode request: %v", err)
		}
	}
	// set default
	if err := validation.SetDefault(po); err != nil {
		u.ApiErrorf("failed to set form default: %v", err)
	}
	// valid
	valid := validation.Validation{}
	if ok, err := valid.Valid(po); err != nil {
		u.ApiErrorf("failed to valid request: %v", err)
	} else if !ok && len(valid.Errors) > 0 {
		u.ApiError(valid.Errors[0].Key+", "+valid.Errors[0].Message, -1000)
	}
}
