package main

import (
	_ "mymallweb/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"fmt"
	"syscall"
)

func init()  {
	beego.Info("开始启动mysql")
	//连接MySQL
	dbUser := beego.AppConfig.String("mysqluser")
	dbPass := beego.AppConfig.String("mysqlpass")
	dbHost := beego.AppConfig.String("mysqlhost")
	dbPort := beego.AppConfig.String("mysqlport")
	dbName := beego.AppConfig.String("mysqldb")

	maxIdleConn, _ := beego.AppConfig.Int("mysql_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"
	//utils.Display("dbLink", dbLink)
	beego.Info(dbLink)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	//设置日志
	fn := "logs/run.log"
	if _, err := os.Stat(fn); err != nil {
		if os.IsNotExist(err) {
			os.Create(fn)
		}
	}
	beego.SetLogger("file", `{"filename":"`+fn+`"}`)
	if beego.BConfig.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
}
func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:

		beego.Info("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		beego.Info("Shutdown gracefully, bye...")
		// do graceful shutdown
	}
	os.Exit(0)
}
func main() {
	beego.Run()
}

