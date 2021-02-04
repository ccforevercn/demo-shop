package models

import (
	"demo-shop/service"
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

// 列表
func (model *Product) List(name string, offset, limit int64) ([]*Product, error) {
	list := []*Product{}
	newOrm := Before()
	_, error := newOrm.QueryTable(&model).Filter("name", name).Filter("is_del", 0).Offset(offset).Limit(limit).All(&list)
	return list, error
}

// 验证编号
func (model *Product) CheckId(id int64)  (int64, error) {
	newOrm := Before()
	return newOrm.QueryTable(&model).Filter("id", id).Filter("is_del", 0).Count()
}

// 插入(单条)
func (model *Product) Insert(insert *Product) (int64, error)  {
	newOrm := Before()
	return newOrm.Insert(insert)
}

// 查询(单条)
func (model *Product) Select(product *Product) error {
	newOrm := Before()
	return newOrm.Read(product)
}

// 修改(单条)
func (model *Product) Update(update *Product) (int64, error) {
	newOrm := Before()
	return newOrm.Update(update, "Name", "Keyword", "Description", "Price", "OldPrice", "Image", "UpdatedAt")
}

// 回收站
func (model *Product) Recycle(id int64) error {
	newOrm := Before()
	_, err :=newOrm.QueryTable(&model).Filter("id", id).Filter("is_del", 0).Update(orm.Params{
		"is_del": 1,
		"updated_at": service.GetStringTime(),
	})
	return err
}

// 删除(前提是在回收站)
func (model *Product) Delete(id int64) error {
	newOrm := Before()
	_, err :=newOrm.QueryTable(&model).Filter("id", id).Filter("is_del", 1).Delete()
	return err
}

func init()  {
	orm.RegisterModel(new(Product))
}
