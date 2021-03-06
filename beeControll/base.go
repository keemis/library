package beeControll

import (
	"github.com/astaxie/beego"
	"github.com/keemis/library/logs"
)

type ControllerItf interface {
	Before() // 执行之前（选择实现）
}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
	Log       logs.Logger
	bodyStore map[string]interface{}
}

// Prepare 执行之前
func (u *BaseController) Prepare() {
	u.Log = logs.New()
	u.logRequest()
	u.parseBody()
	if app, ok := u.AppController.(ControllerItf); ok {
		app.Before()
	}
}

// logRequest 记录请求
func (u *BaseController) logRequest() {
	if u.Ctx == nil || u.Ctx.Input == nil || u.Ctx.Input.Context == nil {
		return
	}
	if u.Ctx.Input.Context.Request.Form == nil {
		_ = u.Ctx.Input.Context.Request.ParseForm()
	}
	u.Log.Debug("Request Method: %v, Path: %v", u.Ctx.Request.Method, u.Ctx.Request.URL.Path)
	u.Log.Debug("Request Form: %v", u.Ctx.Input.Context.Request.Form)
	u.Log.Debug("Request Body: %v", string(u.Ctx.Input.RequestBody))
}
