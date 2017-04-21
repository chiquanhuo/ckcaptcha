package main

import (
	"ckcaptcha/global"
	"ckcaptcha/http"
)

func main() {
	global.Init()
	api := http.NewHTTPAPI()
	api.ServeHTTP()
}
