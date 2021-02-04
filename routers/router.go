package routers

import (
	"demo-shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/product", &controllers.ProductController{})
    beego.Router("/product/details/:id", &controllers.ProductController{}, "get:Details")
    beego.Router("/product/recycle", &controllers.ProductController{}, "put:Recycle")
}
