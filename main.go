package main

import (
	"time"

	"github.com/george518/PPGo_ApiAdmin/models"
	_ "github.com/george518/PPGo_ApiAdmin/routers"

	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/utils"
	cache "github.com/patrickmn/go-cache"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.Run()
}
