package http

import (
	"github.com/gin-gonic/gin"
	"ckcaptcha/global"
	log "github.com/cihub/seelog"
	"net/http"
	"fmt"
	"ckcaptcha/handler"
)


type HTTPAPI struct {
	Engine *gin.Engine
	Handler *handler.HandlerBase
}

// init
func NewHTTPAPI() *HTTPAPI {
	if global.Config.DEVMODE {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(Recovery)
	if global.Config.DEVMODE {
		engine.Use(gin.Logger())
	}

	api := &HTTPAPI{
		Engine: engine,
		Handler: handler.NewHandlerBase(),
	}

	return api
}

func (api *HTTPAPI) ServeHTTP() {
	router := api.Engine
	router.POST("/verify", api.Handler.VerifyHandler)
	router.GET("/req", api.Handler.ReqHandler)
	router.GET("/img", api.Handler.ImgHandler)
	router.GET("/reload", api.Handler.ReloadHandler)

	addr := fmt.Sprintf("%s:%s", global.Config.Listen.Addr, global.Config.Listen.Bind)
	fmt.Println(addr)
	router.Run(addr)
}
func Recovery(c *gin.Context) {
	defer func() {
		for _, item := range c.Errors {
			log.Errorf("Recovery error: ", item.Err.Error())
		}
		if rval := recover(); rval != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			message := fmt.Sprint(rval)
			log.Errorf("Recovery - panic: ", message)
		}
	}()
	c.Next()
}
