package beeControll

import (
	"github.com/astaxie/beego"
	"github.com/keemis/library/logs"
)

type ControllerItf interface {
	Before()
	After()
}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
	log logs.Logger
}

// Prepare 方法之前
func (u *BaseController) Prepare() {
	u.log = logs.New()

	if app, ok := u.AppController.(ControllerItf); ok {
		app.Before()
	}
}

// Finish 方法之后
func (u *BaseController) Finish() {
	if app, ok := u.AppController.(ControllerItf); ok {
		app.After()
	}
}
