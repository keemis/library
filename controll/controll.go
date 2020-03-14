package controll

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ControllerItf interface {
	Before()
}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

// Prepare 初始化请求
func (u *BaseController) Prepare() {

	fmt.Println("00000000000")

	if app, ok := u.AppController.(ControllerItf); ok {
		app.Before()
	}
}

func (u *BaseController) Test() {

	fmt.Println("1111111")

}
