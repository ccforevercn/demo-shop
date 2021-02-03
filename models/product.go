package models

import (
	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id int64 `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Keyword string `json:"keyword" form:"keyword"`
	Description string `json:"description" form:"description"`
	Price float64 `json:"price" form:"price"`
	OldPrice float64 `json:"old_price" form:"old_price"`
	Image string `json:"image" form:"image"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	IsDel int `json:"-" form:"is_del"`
}

func (model *Product) before() orm.Ormer {
	orm.Debug = true
	newOrm := orm.NewOrm()
	newOrm.Using("default")
	return newOrm
}

// 插入
func (model *Product) Insert(data *Product) (int64, error)  {
	newOrm := model.before()
	return newOrm.Insert(data)
}

func init()  {
	orm.RegisterModel(new(Product))
}
