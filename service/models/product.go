package models

import (
	"demo-shop/models"
	"demo-shop/service"
)

type ProductService struct {

}

var (
	timeService = service.TimeService{}
	product = models.Product{}
	errorService = service.ErrorService{}
)

// 列表
func (productService *ProductService) List(name string, page, limit int64) ([]*models.Product, error) {
	offset := service.Paging(page, limit) // 获取起始值
	return product.List(name, offset, limit)
}

// 查看
func (productService *ProductService) Select(selects *models.Product) error {
	return product.Select(selects)
}

// 插入
func (productService *ProductService) Insert(insert *models.Product) (int64, error) {
	insert.CreatedAt = timeService.GetStringTime() // 设置添加时间
	insert.UpdatedAt = timeService.GetStringTime() // 设置修改时间
	return product.Insert(insert) // 添加
}

// 修改
func (productService *ProductService) Update(update *models.Product) *service.ErrorService {
	update.IsDel = 0
	count, err := product.CheckId(update.Id)  // 验证编号
	if count == 0 {
		return errorService.SetError(400, "参数错误")
	}
	if err != nil {
		return errorService.SetError(400, "修改失败")
	}
	update.UpdatedAt = timeService.GetStringTime() // 设置修改时间
	_, err = product.Update(update)
	if err != nil {
		return errorService.SetError(400, "修改失败")
	}
	return nil
}

// 回收站
func (productService *ProductService) Recycle(id int64) *service.ErrorService {
	err := product.Recycle(id)
	if err != nil {
		return errorService.SetError(400, "修改失败")
	}
	return nil
}

// 删除
func (productService *ProductService) Delete(id int64) *service.ErrorService {
	err := product.Delete(id)
	if err != nil {
		return errorService.SetError(400, "删除失败")
	}
	return nil
}
