package controllers

import (
	"demo-shop/models"
	"demo-shop/service"
	serviceModel "demo-shop/service/models"
	"demo-shop/validly"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductController struct {
	beego.Controller
}

// 列表
func (controller *ProductController) Get()  {
	products := []*models.Product{}
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(models.Product{}).All(&products)
	response := &service.Response{}
	data := response.GetSuccess(products, "列表")
	controller.Data["json"] = data
	controller.ServeJSON()
}

// 添加
func (controller *ProductController) Post()  {
	product := models.Product{}
	response := service.Response{}
	productValidly := validly.ProductValidly{}
	controller.ParseForm(&product)
	err := productValidly.Insert(&product)
	if err != nil {
		message, _ := err.Error()
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	product.Id = 0
	productService := serviceModel.ProductService{}
	id, error := productService.Insert(&product)
	if error != nil {
		controller.Data["json"] = response.GetError(error.Error())
		controller.ServeJSON()
		return
	}
	product.Id = id
	controller.Data["json"] = response.GetSuccess(product, "添加成功")
	controller.ServeJSON()
}