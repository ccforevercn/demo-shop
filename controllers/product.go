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

var (
	response = service.Response{}
	productService = serviceModel.ProductService{}
	productValidly = validly.ProductValidly{}
)

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
	controller.ParseForm(&product)  // 获取参数
	product.Id = 0  // 重置编号为0
	err := productValidly.Insert(&product) // 验证参数
	if err != nil {
		message, _ := err.Error() // 参数错误提示
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	id, error := productService.Insert(&product) // 添加
	if error != nil {
		controller.Data["json"] = response.GetError(error.Error())  // 添加失败错误提示
		controller.ServeJSON()
		return
	}
	product.Id = id // 重置编号为数据库编号
	controller.Data["json"] = response.GetSuccess(product, "添加成功")
	controller.ServeJSON()
}

// 修改
func (controller ProductController) Put()  {
	product := models.Product{}
	controller.ParseForm(&product)
	err := productValidly.Update(&product)
	if err != nil {
		message, _ := err.Error()
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	productService := serviceModel.ProductService{}
	err = productService.Update(&product)
	if err != nil {
		message, _ := err.Error()
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	controller.Data["json"] = response.GetSuccess(product, "修改成功")
	controller.ServeJSON()
}