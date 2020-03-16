package beeControll

import (
	"github.com/astaxie/beego"
	"github.com/keemis/library/logs"
)

type ControllerItf interface {
	Before() // 执行之前（选择实现）
	After()  // 执行之后（选择实现）
}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
	log       logs.Logger
	bodyStore map[string]interface{}
}

// Prepare 执行之前
func (u *BaseController) Prepare() {
	u.log = logs.New()
	u.logRequest()
	u.resolveBody()
	if app, ok := u.AppController.(ControllerItf); ok {
		app.Before()
	}
}

// Finish 执行之后
func (u *BaseController) Finish() {
	if app, ok := u.AppController.(ControllerItf); ok {
		app.After()
	}
}

// logRequest 记录请求
func (u *BaseController) logRequest() {
	if u.Ctx.Input.Context.Request.Form == nil {
		u.Ctx.Input.Context.Request.ParseForm()
	}
	u.log.Debug("request form: ", u.Ctx.Input.Context.Request.Form)
	u.log.Debug("request body: ", string(u.Ctx.Input.RequestBody))
}
