package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var o orm.Ormer

//数据库连接
func Connect() {
	dbUser := beego.AppConfig.String("mysqluser")
	dbPass := beego.AppConfig.String("mysqlpass")
	dbHost := beego.AppConfig.String("mysqlhost")
	dbPort := beego.AppConfig.String("mysqlport")
	dbName := beego.AppConfig.String("mysqldb")
	if beego.BConfig.RunMode == "docker" {
		if os.Getenv("MYSQL_USER") != "" {
			dbUser = os.Getenv("MYSQL_USER")
		}
		if os.Getenv("MYSQL_PASS") != "" {
			dbPass = os.Getenv("MYSQL_PASS")
		}
		if os.Getenv("MYSQL_HOST") != "" {
			dbHost = os.Getenv("MYSQL_HOST")
		}
		if os.Getenv("MYSQL_PORT") != "" {
			dbPort = os.Getenv("MYSQL_PORT")
		}
		if os.Getenv("MYSQL_DB") != "" {
			dbName = os.Getenv("MYSQL_DB")
		}
	}

	maxIdleConn, _ := beego.AppConfig.Int("mysql_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"
	//utils.Display("dbLink", dbLink)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("数据库连接错误:", err)
		os.Exit(2)
		return
	}
	err = orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)
	orm.Debug = true
	if err != nil {
		beego.Error("数据库连接错误:", err)
		os.Exit(2)
		return
	}
}
