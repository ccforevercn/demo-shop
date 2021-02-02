package main

import (
	_ "demo-shop/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	db "demo-shop/service/database"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		fmt.Println(err.Error())
	}
	database, _ := db.NewDataBase(conf.String("db::dbType"))
	orm.RegisterDriver(database.GetDriverName(), database.GetDriver())
	orm.RegisterDataBase(database.GetAliasName(), database.GetDriverName(), database.GetStr())
}

func main() {
	beego.Run()
}

