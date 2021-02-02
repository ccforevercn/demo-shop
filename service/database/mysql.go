package database

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
)

type MysqlDataBase struct {
}

func NewMysqlDataBase() DataBase {
	database := MysqlDataBase{}
	return &database
}
// 获取
func (database *MysqlDataBase) GetStr() string{
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		fmt.Println(err.Error())
	}

	dbUser := conf.String("db::dbUser")
	dbPass := conf.String("db::dbPass")
	dbHost := conf.String("db::dbHost")
	dbPort := conf.String("db::dbPort")
	dbName := conf.String("db::dbName")
	dbCharset := conf.String("db::dbCharset")
	dbLocation := conf.String("db::dbLocation")
	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName+ "?charset=" + dbCharset + "&loc=" + dbLocation

}

func (database *MysqlDataBase) GetDriverName() string{
	return "mysql"
}

func (database *MysqlDataBase) GetAliasName() string{
	return "default"
}

func (database *MysqlDataBase) GetDriver() orm.DriverType{
	return orm.DRMySQL
}


func init(){
	Register("mysql",NewMysqlDataBase)
}


