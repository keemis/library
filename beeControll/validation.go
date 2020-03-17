package beeControll

import (
	"encoding/json"
	"reflect"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/pkg/errors"

	"github.com/keemis/library/validation"
)

// ValidForm 校验、解析Query、Form、Body到Po
func (u *BaseController) ValidForm(po interface{}) {
	rv := reflect.ValueOf(po)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		u.ApiError("param must be ptr")
	}
	// parse
	byts, err := u.parseFormBody()
	if err != nil {
		u.ApiError(err.Error())
	}
	// use jsoniter ext
	if len(byts) > 2 {
		extra.RegisterFuzzyDecoders()
		extra.RegisterTimeAsInt64Codec(time.Second)
		if err := jsoniter.Unmarshal(byts, po); err != nil {
			u.ApiErrorf("failed to decode request: %v", err)
		}
	}
	// valid
	valid := validation.Validation{}
	if ok, err := valid.Valid(po); err != nil {
		u.ApiErrorf("failed to valid request: %v", err)
	} else if !ok && len(valid.Errors) > 0 {
		u.ApiError(valid.Errors[0].Key + ", " + valid.Errors[0].Message)
	}
}

// parseFormBody 解析Query、Form、Body
func (u *BaseController) parseFormBody() ([]byte, error) {
	// ParseForm
	if u.Ctx.Input.Context.Request.Form == nil {
		_ = u.Ctx.Input.Context.Request.ParseForm()
	}
	// ParseBody
	bodyMap := make(map[string]interface{})
	if len(u.Ctx.Input.RequestBody) > 2 {
		_ = jsoniter.Unmarshal(u.Ctx.Input.RequestBody, &bodyMap)
	}
	// form + body
	tmp := make(map[string]interface{})
	for k, vs := range u.Ctx.Input.Context.Request.Form {
		v := ""
		if len(vs) > 0 {
			v = vs[0]
		}
		tmp[k] = v
	}
	for k, v := range bodyMap {
		tmp[k] = v
	}
	byts, err := json.Marshal(tmp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode form")
	}
	return byts, nil
}
