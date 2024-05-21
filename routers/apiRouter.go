package routers

import (
	"B2CProject/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSRouter("/", &api.V1Controller{}),
		beego.NSRouter("/menu", &api.V1Controller{}, "get:Menu"),
	)
	beego.AddNamespace(ns)
}
