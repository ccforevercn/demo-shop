package controllers

import (
	serviceModel "demo-shop/service/models"
	"demo-shop/service/response"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

var (
	categoryService = serviceModel.CategoryService{}
	//categoryValidly = validly.CategoryValidly{}
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
