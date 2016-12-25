package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlStr := beego.AppConfig.String("mysql")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlStr)
	orm.RegisterModel(new(Article))
	orm.RunSyncdb("default", false, true)

}