package beeControll

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

// parseBody 解析application/json参数
func (u *BaseController) parseBody() {
	u.bodyStore = make(map[string]interface{})
	// Parse Form
	if u.Ctx.Input.Context.Request.Form == nil {
		_ = u.Ctx.Input.Context.Request.ParseForm()
	}
	for k, vs := range u.Ctx.Input.Context.Request.Form {
		v := ""
		if len(vs) > 0 {
			v = vs[0]
		}
		u.bodyStore[k] = v
	}
	// Parse Body
	if u.Ctx.Request.Method == "POST" && len(u.Ctx.Input.RequestBody) > 2 {
		bodyMap := make(map[string]interface{})
		if len(u.Ctx.Input.RequestBody) > 2 {
			_ = jsoniter.Unmarshal(u.Ctx.Input.RequestBody, &bodyMap)
		}
		for k, v := range bodyMap {
			u.bodyStore[k] = v
		}
	}
}

// GetBodyString 获取application/json参数
// 用法 & 作用 ：类同beego的 GetString()
func (u *BaseController) GetBodyString(key string, def ...string) string {
	res := ""
	if len(def) > 0 {
		res = def[0]
	}
	if val, ok := u.bodyStore[key].(string); ok {
		return val
	}
	return res
}

// GetBodyInt 获取application/json参数
// 用法 & 作用 ：类同beego的 GetBodyInt()
func (u *BaseController) GetBodyInt(key string, def ...int) int {
	res := 0
	if len(def) > 0 {
		res = def[0]
	}
	v, ok := u.bodyStore[key]
	if !ok {
		return res
	}
	switch data := v.(type) {
	case string:
		res, _ = strconv.Atoi(data)
	case float64:
		res = int(data)
	case int:
		res = data
	}
	return res
}

// GetBodyInt64 获取application/json参数
// 用法 & 作用 ：类同beego的 GetInt64()
func (u *BaseController) GetBodyInt64(key string, def ...int64) int64 {
	res := int64(0)
	if len(def) > 0 {
		res = def[0]
	}
	v, ok := u.bodyStore[key]
	if !ok {
		return res
	}
	switch data := v.(type) {
	case string:
		res, _ = strconv.ParseInt(data, 10, 64)
	case float64:
		res = int64(data)
	case int:
		res = int64(data)
	case int64:
		res = data
	}
	return res
}

// GetBodyFloat64 获取application/json参数
// 用法 & 作用 ：类同beego的 GetInt64()
func (u *BaseController) GetBodyFloat64(key string, def ...float64) float64 {
	var res float64
	if len(def) > 0 {
		res = def[0]
	}
	v, ok := u.bodyStore[key]
	if !ok {
		return res
	}
	switch data := v.(type) {
	case string:
		res, _ = strconv.ParseFloat(data, 64)
	case float64:
		res = data
	case float32:
		res = float64(data)
	case int:
		res = float64(data)
	case int64:
		res = float64(data)
	}
	return res
}

// GetBodyBool 获取application/json参数
// 用法 & 作用 ：类同beego的 GetBodyBool()
func (u *BaseController) GetBodyBool(key string, def ...bool) bool {
	res := false
	if len(def) > 0 {
		res = def[0]
	}
	v, ok := u.bodyStore[key]
	if !ok {
		return res
	}
	switch data := v.(type) {
	case bool:
		res = data
	case int:
		if data == 0 {
			res = false
		} else {
			res = true
		}
	}
	return res
}

// GetBodyBool 获取application/json参数
// 用法 & 作用 ：类同beego的 GetBodyItf()
func (u *BaseController) GetBodyItf(key string) (interface{}, bool) {
	itf, ok := u.bodyStore[key]
	return itf, ok
}
