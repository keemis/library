package beeControll

import (
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/keemis/library/errs"
)

// apiResult 接口返回结构
type apiResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Option 设置选项
type Option func(*apiResult)

// OptCode 设置Code
func OptCode(code int) Option {
	return func(o *apiResult) {
		o.Code = code
	}
}

// OptMsg 设置Msg
func OptMsg(msg string) Option {
	return func(o *apiResult) {
		o.Msg = msg
	}
}

// OptMsg 设置Data
func OptData(data interface{}) Option {
	return func(o *apiResult) {
		o.Data = data
	}
}

// apiSuccess RPC返回正确
func (u *BaseController) ApiSuccess(data interface{}) {
	u.output(apiResult{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// apiError RPC返回错误
func (u *BaseController) ApiError(msg string) {
	u.output(apiResult{
		Code: -1000,
		Msg:  msg,
		Data: nil,
	})
}

// apiErrorf RPC返回错误
func (u *BaseController) ApiErrorf(format string, a ...interface{}) {
	u.ApiError(fmt.Sprintf(format, a))
}

// apiResult RPC返回
func (u *BaseController) ApiResult(opt ...Option) {
	res := apiResult{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
	for _, o := range opt {
		o(&res)
	}
	u.output(res)
}

// apiStruct RPC返回Error (github.com/keemis/library/errs)
func (u *BaseController) ApiStruct(err error) {
	u.output(apiResult{
		Code: errs.GetCode(err),
		Msg:  errs.GetMsg(err),
		Data: errs.GetData(err),
	})
}

// output 输出结果
func (u *BaseController) output(data apiResult) {
	hasIndent := true
	if beego.BConfig.RunMode == beego.PROD {
		hasIndent = false
	}
	var content []byte
	var err error
	// use jsoniter ext
	extra.RegisterFuzzyDecoders()
	extra.RegisterTimeAsInt64Codec(time.Second)
	if hasIndent {
		content, err = jsoniter.MarshalIndent(data, "", "  ")
	} else {
		content, err = jsoniter.Marshal(data)
	}
	if err != nil {
		http.Error(u.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	// log
	if data.Code != 0 {
		u.Log.Warn("rpc response: %s", string(content))
	}
	// output
	u.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	u.Ctx.Output.Header("Server", "xService")
	u.Ctx.Output.Body(content)
	u.StopRun()
}
