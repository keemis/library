package session

import (
	"github.com/astaxie/beego/session"
)

// GlobalSession 全局Session
var GlobalSession *session.Manager

// Init 初始化Session
func Init() {
	gs, err := session.NewManager("memory", &session.ManagerConfig{
		CookieName:      "sessid",
		EnableSetCookie: true,
		Gclifetime:      1800,
		Maxlifetime:     3600,
		CookieLifeTime:  3600,
		Secure:          false,
		ProviderConfig:  "",
	})
	if err != nil {
		panic(err.Error())
	}
	go gs.GC()
	GlobalSession = gs
}
