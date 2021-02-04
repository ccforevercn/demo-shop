package models

import (
	"demo-shop/models"
	"demo-shop/service"
)

type CategoryService struct {

}
var (
	category = models.Category{}
)

// 列表
func (categoryService CategoryService) List(name string, cid, page, limit int64) ([]*models.Category, error) {
	offset := service.Paging(page, limit) // 获取起始值
	return category.List(name, cid, offset, limit)
}

// 添加
func (categoryService CategoryService) Insert(insert *models.Category) (int64, error) {
	insert.CreatedAt = service.GetStringTime() // 设置添加时间
	insert.UpdatedAt = service.GetStringTime() // 设置修改时间
	return category.Insert(insert)
}
