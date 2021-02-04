package routers

import (
	"demo-shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
	router := beego.NewNamespace("/v1",
		beego.NSRouter("/", &controllers.MainController{}),
		beego.NSNamespace("/admin",
			beego.NSRouter("/product", &controllers.ProductController{}),
			beego.NSRouter("/product/details/:id", &controllers.ProductController{}, "get:Details"),
			beego.NSRouter("/product/recycle", &controllers.ProductController{}, "put:Recycle"),
		),
		beego.NSNamespace("/api",
			beego.NSRouter("/", &controllers.MainController{}),
		),
	)
	beego.AddNamespace(router)
}
