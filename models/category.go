package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id int64 `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Cid int64 `json:"cid" form:"cid"`
	Image string `json:"image" form:"image"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	IsDel int64 `json:"-" form:"-"`
}

func init()  {
	orm.RegisterModel(new(Category))
}

// 列表
func (model Category) List(name string, cid, offset, limit int64) ([]*Category, error) {
	list := []*Category{}
	newOrm := Before()
	_, error := newOrm.QueryTable(&model).Filter("name", name).Filter("cid", cid).Filter("is_del", 0).Offset(offset).Limit(limit).All(&list)
	return list, error
}

// 添加
func (model Category) Insert(insert *Category) (int64, error) {
	newOrm := Before()
	return newOrm.Insert(insert)
}