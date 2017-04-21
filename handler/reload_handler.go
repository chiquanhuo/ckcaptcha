package handler

import (
	"github.com/gin-gonic/gin"
)

func (handler *HandlerBase) ReloadHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id := c.Query("id")
	res := handler.CaptchaInfo.ReloadCaptcha(id)
	if res == false {
		c.JSON(200, gin.H{
			"c": -1,
			"msg": "reload error",
		})
		return
	}

	c.JSON(200, gin.H{
		"c": 0,
		"err": "",
	})
	return
}