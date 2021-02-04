package main

import (
	_ "sampleapp/routers"
	beego "github.com/beego/beego/v2/server/web"
	 "github.com/beego/beego/v2/server/web"
  _ "github.com/beego/beego/v2/server/web/session/redis"
)



func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "redis"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "10.156.0.2:6379,100,69000ca0-56ef-4600-a8c4-9b7f3176f0f0"
}


func main() {
	beego.Run()
}

