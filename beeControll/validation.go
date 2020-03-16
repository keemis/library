package beeControll

import (
	json "github.com/json-iterator/go"
	"reflect"

	"github.com/keemis/library/validation"
)

// ValidQuery 校验、解析Query到Po
func (u *BaseController) ValidQuery(po interface{}) {
	rv := reflect.ValueOf(po)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		u.ApiError("ValidQuery func params po must be ptr")
	}
	if u.Ctx.Input.Context.Request.Form == nil {
		if err := u.Ctx.Input.Context.Request.ParseForm(); err != nil {
			u.ApiErrorf("failed to parse form: %v", err)
		}
	}
	tmp := make(map[string]string)
	for k, vs := range u.Ctx.Input.Context.Request.Form {
		v := ""
		if len(vs) > 0 {
			v = vs[0]
		}
		tmp[k] = v
	}
	byt, err := json.Marshal(tmp)
	if err != nil {
		u.ApiErrorf("failed to encode query: %v", err)
	}
	if len(byt) <= 2 {
		return
	}
	if err := json.Unmarshal(byt, po); err != nil {
		u.ApiErrorf("failed to decode request: %v", err)
	}
	valid := validation.Validation{}
	if ok, err := valid.Valid(po); err != nil {
		u.ApiErrorf("failed to valid request: %v", err)
	} else if !ok && len(valid.Errors) > 0 {
		u.ApiError(valid.Errors[0].Key + ", " + valid.Errors[0].Message)
	}
}

// ValidBody 校验、解析Body到Po
func (u *BaseController) ValidBody(po interface{}) {
	rv := reflect.ValueOf(po)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		u.ApiError("validBody func params po must be ptr")
	}
	body := u.Ctx.Input.RequestBody
	if len(body) <= 2 {
		return
	}
	if err := json.Unmarshal(body, po); err != nil {
		u.ApiErrorf("failed to decode request: %v", err)
	}
	valid := validation.Validation{}
	if ok, err := valid.Valid(po); err != nil {
		u.ApiErrorf("failed to valid request: %v", err)
	} else if !ok && len(valid.Errors) > 0 {
		u.ApiError(valid.Errors[0].Key + ", " + valid.Errors[0].Message)
	}
}
