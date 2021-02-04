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

type CategoryController struct {
	beego.Controller
}

var (
	categoryService = serviceModel.CategoryService{}
	categoryValidly = validly.CategoryValidly{}
)

// 列表
func (controller *CategoryController) Get()  {
	name := fmt.Sprint(controller.GetString("name"))
	paramCid := fmt.Sprint(controller.GetString("cid"))
	cid, _ := strconv.ParseInt(paramCid, 10, 64)
	paramPage := fmt.Sprint(controller.GetString("page"))
	page, _ := strconv.ParseInt(paramPage, 10, 64)
	paramLimit := fmt.Sprint(controller.GetString("limit"))
	limit, _ := strconv.ParseInt(paramLimit, 10, 64)
	category, error := categoryService.List(name, cid, page, limit)
	if error != nil {
		controller.Data["json"] = response.GetError(error.Error())
		controller.ServeJSON()
		return
	}
	data := response.GetSuccess(category, "列表")
	controller.Data["json"] = data
	controller.ServeJSON()
}

// 添加
func (controller CategoryController) Post()  {
	category := models.Category{}
	controller.ParseForm(&category)  // 获取参数
	category.Id = 0  // 重置编号为0
	err := categoryValidly.Insert(&category) // 验证参数
	if err != nil {
		message, _ := err.Error() // 参数错误提示
		controller.Data["json"] = response.GetError(message)
		controller.ServeJSON()
		return
	}
	id, error := categoryService.Insert(&category) // 添加
	if error != nil {
		controller.Data["json"] = response.GetError(error.Error())  // 添加失败错误提示
		controller.ServeJSON()
		return
	}
	category.Id = id // 重置编号为数据库编号
	controller.Data["json"] = response.GetSuccessMessage("添加成功")
	controller.ServeJSON()
}
