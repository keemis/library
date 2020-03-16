package beeControll

import (
	"encoding/json"

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
func (u *BaseController) apiSuccess(data interface{}) {
	u.apiOutput(apiResult{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// apiError RPC返回错误
func (u *BaseController) apiError(data interface{}) {
	u.apiOutput(apiResult{
		Code: -1000,
		Msg:  "error",
		Data: data,
	})
}

// apiResult RPC返回
func (u *BaseController) apiResult(opt ...Option) {
	res := apiResult{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
	for _, o := range opt {
		o(&res)
	}
	u.apiOutput(res)
}

// apiStruct RPC返回Error (github.com/keemis/library/errs)
func (u *BaseController) apiStruct(err error) {
	u.apiOutput(apiResult{
		Code: errs.GetCode(err),
		Msg:  errs.GetMsg(err),
		Data: errs.GetData(err),
	})
}

// apiOutput 返回结果
func (u *BaseController) apiOutput(data apiResult) {
	if data.Code != 0 {
		if byt, err := json.Marshal(data); err == nil {
			u.log.Warn("rpc response: %s", string(byt))
		}
	}
	u.Ctx.Output.Header("Server", "xService")
	u.Data["json"] = data
	u.ServeJSON()
	u.StopRun()
}
