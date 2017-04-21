package handler

import (
	"ckcaptcha/code"
	"ckcaptcha/global"
	"time"
)

type HandlerBase struct {
	CaptchaInfo *code.CaptchaInfo
}

func NewHandlerBase() *HandlerBase {
	Handler := &HandlerBase{
		CaptchaInfo : code.NewCpatchaInfo(global.Config.Code.Len, time.Duration(global.Config.Code.Expire)*time.Minute, global.Config.Code.Width, global.Config.Code.Height),
	}
	return Handler
}