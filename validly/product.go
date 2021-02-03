package validly

import (
	"demo-shop/models"
	"demo-shop/service"
	"github.com/astaxie/beego/validation"
)

type ProductValidly struct {

}

var (
	valid = validation.Validation{}
	errorService = service.ErrorService{}
)

// 插入验证
func (productValidly *ProductValidly) Insert(product *models.Product) *service.ErrorService {
	valid.Required(product.Name, "name").Message("请填写名称")
	valid.MaxSize(product.Name, 60, "name").Message("名称不能超过60个汉字")
	valid.MaxSize(product.Keyword, 80, "keyword").Message("关键字不能超过80个汉字")
	valid.MaxSize(product.Description, 100, "description").Message("关键字不能超过100个汉字")
	valid.Required(product.Image, "image").Message("请选择图片")
	valid.Range(product.IsDel, 0, 1, "is_del").Message("参数错误")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errorService.SetError(400, err.Error())
		}
	}
	if product.Price <= 0 {
		return errorService.SetError(400, "价格不能小于0元")
	}
	if product.OldPrice <= 0 {
		return errorService.SetError(400, "原价不能小于0元")
	}
	if len(product.Keyword) == 0 {
		product.Keyword = product.Name
	}
	if len(product.Description) == 0 {
		product.Description = product.Name
	}
	return  nil
}

// 修改
func (productValidly ProductValidly) Update(product *models.Product) *service.ErrorService {
	valid.Required(product.Name, "name").Message("请填写名称")
	valid.MaxSize(product.Name, 60, "name").Message("名称不能超过60个汉字")
	valid.MaxSize(product.Keyword, 80, "keyword").Message("关键字不能超过80个汉字")
	valid.MaxSize(product.Description, 100, "description").Message("关键字不能超过100个汉字")
	valid.Required(product.Image, "image").Message("请选择图片")
	valid.Range(product.IsDel, 0, 1, "is_del").Message("参数错误")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errorService.SetError(400, err.Error())
		}
	}
	if product.Id <= 0 {
		return errorService.SetError(400, "参数错误")
	}
	if product.Price <= 0 {
		return errorService.SetError(400, "价格不能小于0元")
	}
	if product.OldPrice <= 0 {
		return errorService.SetError(400, "原价不能小于0元")
	}
	if len(product.Keyword) == 0 {
		product.Keyword = product.Name
	}
	if len(product.Description) == 0 {
		product.Description = product.Name
	}
	return  nil
}