package models

import "github.com/astaxie/beego/orm"

func Before() orm.Ormer {
	orm.Debug = true
	newOrm := orm.NewOrm()
	newOrm.Using("default")
	return newOrm
}