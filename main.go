package main

import (
	"PPGo_ApiAdmin/models"
	_ "PPGo_ApiAdmin/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
