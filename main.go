package main

import (
	"wechat-mall-backend/web"
)

func main() {
	app := &web.App{}
	app.Initialize()
	app.Run()
}
