package controllers

import (
	"demo-shop/models"
	serviceModel "demo-shop/service/models"
	"demo-shop/service/response"
	"demo-shop/validly"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ProductController struct {
	beego.Controller
}

var (
	productService = serviceModel.ProductService{}
	productValidly = validly.ProductValidly{}
)

// 列表
func (controller *ProductController) Get()  {
	name := fmt.Sprint(controller.GetString("name"))
	paramPage := fmt.Sprint(controller.GetString("page"))
	page, _ := strconv.ParseInt(paramPage, 10, 64)
	paramLimit := fmt.Sprint(controller.GetString("limit"))
	limit, _ := strconv.ParseInt(paramLimit, 10, 64)
	products, error := productService.List(name, page, limit)
	if error != nil {
		controller.Data["json"] = response.GetError(error.Error())
		controller.ServeJSON()
		return
	}
	data := response.GetSuccess(products, "列表")
	controller.Data["json"] = data
	controller.ServeJSON()
}

// 详情
func (controller *ProductController) Details() {
	string := fmt.Sprint(controller.Ctx.Input.Param(":id"))
	id, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		controller.Data["json"] = response.GetError("参数错误")  // 添加错误提示
		controller.ServeJSON()
		return
	}
	product := models.Product{}
	if id < 1 {
		controller.Data["json"] = response.GetError("参数错误")  // 添加错误提示
		controller.ServeJSON()
		return
	}
	product.Id = id // 设置编号
	error := productService.Select(&product) // 添加
	if error != nil {
		controller.Data["json"] = response.GetError(error.Error())
		controller.ServeJSON()
		return
	}
	controller.Data["json"] = response.GetSuccess(product, "详情")
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
	controller.Data["json"] = response.GetSuccessMessage("添加成功")
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
	controller.Data["json"] = response.GetSuccessMessage("修改成功")
	controller.ServeJSON()
}

// 回收站
func (controller *ProductController) Recycle()  {
	paramId := controller.Ctx.Request.PostForm.Get("id")
	id, _  := strconv.ParseInt(paramId, 10, 64)
	err := productService.Recycle(id)
	if err != nil {
		message, _ := err.Error()
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	controller.Data["json"] = response.GetSuccessMessage("添加成功")
	controller.ServeJSON()
}

// 删除
func (controller *ProductController) Delete()  {
	paramId := controller.Ctx.Request.PostForm.Get("id")
	id, _  := strconv.ParseInt(paramId, 10, 64)
	err := productService.Delete(id)
	if err != nil {
		message, _ := err.Error()
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	controller.Data["json"] = response.GetSuccessMessage("删除成功")
	controller.ServeJSON()
}