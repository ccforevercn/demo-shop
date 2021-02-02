package controllers

import (
	"demo-shop/models"
	"demo-shop/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductController struct {
	beego.Controller
}


func (controller *ProductController) List()  {
	products := []*models.Product{}
	o := orm.NewOrm()
	o.Using("default")
	o.QueryTable(models.Product{}).All(&products)
	response := &service.Response{}
	data := response.GetSuccess(products, "列表")
	controller.Data["json"] = data
	controller.ServeJSON()
}
