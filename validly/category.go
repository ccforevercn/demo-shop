package validly

import (
	"demo-shop/models"
	"demo-shop/service"
)

type CategoryValidly struct {

}

func (categoryValidly CategoryValidly) Insert(category *models.Category) *service.ErrorService {
	return nil
}

