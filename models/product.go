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

// 验证编号
func (model *Product) CheckId(id int64)  (int64, error) {
	newOrm := model.before()
	return newOrm.QueryTable(&model).Filter("id", id).Filter("is_del", 0).Count()
}

// 插入(单条)
func (model *Product) Insert(insert *Product) (int64, error)  {
	newOrm := model.before()
	return newOrm.Insert(insert)
}

// 查询(单条)
func (model *Product) Select(product *Product) error {
	newOrm := model.before()
	return newOrm.Read(product)
}

// 修改(单条)
func (model *Product) Update(update *Product) (int64, error) {
	newOrm := model.before()
	return newOrm.Update(update, "Name", "Keyword", "Description", "Price", "OldPrice", "Image", "UpdatedAt")
}

func init()  {
	orm.RegisterModel(new(Product))
}
