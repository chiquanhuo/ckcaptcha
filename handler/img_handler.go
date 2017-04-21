package handler

import(
	"github.com/gin-gonic/gin"
	"strconv"
	"bytes"
	log "github.com/cihub/seelog"
)

func (handler *HandlerBase) ImgHandler(c *gin.Context) {
	id := c.Query("id")
	w := c.Query("w")
	width, _ := strconv.Atoi(w)

	h := c.Query("h")
	height, _ := strconv.Atoi(h)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	var buff bytes.Buffer
	img, err := handler.CaptchaInfo.GetCaptchaImage(buff, id, width, height)
	if img == nil {
		c.Data(200, "image/png", nil)
	} else {
		c.Data(200, "image/png", img)
	}

	if err != nil {
		log.Errorf("image error - %s", err.Error())
	}

	return
}