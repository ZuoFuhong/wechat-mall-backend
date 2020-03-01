package main

import (
	"wechat-mall-web/web"
)

func main() {
	app := &web.App{}
	app.Initialize()
	app.Run("127.0.0.1:8080")
}
