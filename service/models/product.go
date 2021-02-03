package models

import (
	"demo-shop/models"
	"demo-shop/service"
)

type ProductService struct {

}

func (productService *ProductService) List()  {

}

// 插入
func (productService *ProductService) Insert(data *models.Product) (int64, error) {
	timeService := service.TimeService{}
	product := models.Product{}
	data.CreatedAt = timeService.GetStringTime()
	data.UpdatedAt = timeService.GetStringTime()
	return product.Insert(data)
}
