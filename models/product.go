package models

import "github.com/astaxie/beego/orm"

type Product struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Keyword string `json:"keyword"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	OldPrice float64 `json:"old_price"`
	Image string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDel int `json:"is_del"`
}

func init()  {
	orm.RegisterModel(new(Product))
}
