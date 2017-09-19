package main

import (
	_ "MaintainServer/routers"
	"github.com/astaxie/beego"
	"MaintainServer/models"
)

func main() {
	models.InitDB()
	models.InitService()
	beego.SetStaticPath("/static","static")
	beego.Run()
}

