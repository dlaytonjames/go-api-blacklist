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

func RegisterMasterDB() {
	dbMasterHost := beego.AppConfig.String("dbMasterHost")
	dbMasterPort := beego.AppConfig.String("dbMasterPort")
	dbMasterSchema := beego.AppConfig.String("dbMasterSchema")
	dbMasterUser := beego.AppConfig.String("dbMasterUser")
	dbMasterPasswd := beego.AppConfig.String("dbMasterPasswd")

	jdbcMasterUrl := dbMasterUser + ":" + dbMasterPasswd + "@tcp(" + dbMasterHost + ":" + dbMasterPort + ")/" + dbMasterSchema + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to mysql master server %v successfully !", dbMasterHost))
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	orm.RegisterDataBase("default", _DB_Driver, jdbcMasterUrl, 30)
}

func RegisterSlaveDB() {
	dbSlaveHost := beego.AppConfig.String("dbSlaveHost")
	dbSlavePort := beego.AppConfig.String("dbSlavePort")
	dbSlaveSchema := beego.AppConfig.String("dbSlaveSchema")
	dbSlaveUser := beego.AppConfig.String("dbSlaveUser")
	dbSlavePasswd := beego.AppConfig.String("dbSlavePasswd")

	jdbcSlaveUrl := dbSlaveUser + ":" + dbSlavePasswd + "@tcp(" + dbSlaveHost + ":" + dbSlavePort + ")/" + dbSlaveSchema + "?charset=utf8"
	beego.Info(fmt.Sprintf("connect to mysql slave server %v successfully !", dbSlaveHost))
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	orm.RegisterDataBase("slave", _DB_Driver, jdbcSlaveUrl, 30)
}
