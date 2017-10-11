package main

import (
	"github.com/george518/PPGo_ApiAdmin/models"
	_ "github.com/george518/PPGo_ApiAdmin/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
