package beeControll

import (
	"encoding/json"
	"reflect"

	"github.com/keemis/library/validation"
)

// ValidBody 校验、解析Body到Po
func (u *BaseController) ValidBody(po interface{}) {
	rv := reflect.ValueOf(po)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		u.ApiError("validBody func params po must be ptr")
	}
	body := u.Ctx.Input.RequestBody
	if len(body) <= 2 {
		body = []byte("{}")
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
