package handler

import (
	"github.com/gin-gonic/gin"
)

func (handler *HandlerBase) VerifyHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, _ := c.GetPostForm("id")
	input, _ := c.GetPostForm("input")
	res := handler.CaptchaInfo.VerifyCaptcha(id, input)
	if !res {
		c.JSON(200, gin.H{
			"c": -2,
			"err": "verify error",
		})
		return
	}

	c.JSON(200, gin.H{
		"c": 0,
		"msg": "verify success",
	})
	return
}