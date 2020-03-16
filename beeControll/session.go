package beeControll

import (
	beeSess "github.com/astaxie/beego/session"
	"github.com/keemis/library/session"
)

// SessionStart 开启beego session
func (u *BaseController) SessionStart() beeSess.Store {
	if session.GlobalSession == nil {
		u.apiErrorf("please initialize session first")
	}
	sess, err := session.GlobalSession.SessionStart(u.Ctx.ResponseWriter, u.Ctx.Request)
	if err != nil {
		u.apiErrorf("session start error: %v", err)
	}
	defer sess.SessionRelease(u.Ctx.ResponseWriter)
	return sess
}
