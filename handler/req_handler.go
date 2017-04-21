package handler

import (
	"github.com/gin-gonic/gin"
)

func (handler *HandlerBase) ReqHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id := handler.CaptchaInfo.GetCaptchaId()

	c.JSON(200, gin.H{
		"c": 0,
		"err": "",
		"data": id,
	})
	return
}