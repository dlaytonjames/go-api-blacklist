package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_Driver = "mysql"
)

func RegisterDB() {
	db_host := beego.AppConfig.String("dbHost")
	db_port := beego.AppConfig.String("dbPort")
	db_schema := beego.AppConfig.String("dbSchema")
	db_user := beego.AppConfig.String("dbUser")
	db_passwd := beego.AppConfig.String("dbPasswd")

	jdbcUrl := db_user + ":" + db_passwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_schema + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to mysql server %v successfully !", db_host))
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	orm.RegisterDataBase("default", _DB_Driver, jdbcUrl, 30)

}
