package database

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/honorjoey/gin-xorm/log"
	"time"
	"xorm.io/core"
)
// mysql
var x *xorm.Engine
//var dbLogger core.ILogger

type XLogger struct {
	log.Log
}

var xLog XLogger

func init() {
	xLog.Tag = "MySQL"

	var err error
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		xLog.Error(err)
	}

	username := cfg.Section("mysql").Key("username").Value()
	password := cfg.Section("mysql").Key("password").Value()
	host := cfg.Section("mysql").Key("host").Value()
	port := cfg.Section("mysql").Key("port").Value()
	database := cfg.Section("mysql").Key("database").Value()

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, database)

	x, err = xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		xLog.Error(err)
	}
	//defer engine.Close()

	x.ShowSQL(true)

	//设置连接池的空闲数大小
	x.SetMaxIdleConns(5)
	//设置最大打开连接数
	x.SetMaxOpenConns(5)

	x.SetMapper(core.SnakeMapper{})

	//dbLogger := x.Logger()

	if err = x.Ping(); err != nil {
		xLog.Error(err)
	}

	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				xLog.Error("database 	connect error: %#v\n", err.Error())
			}
		}
	}(x)
}
