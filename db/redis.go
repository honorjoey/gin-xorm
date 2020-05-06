package db

import (
	"github.com/go-ini/ini"
	"github.com/gomodule/redigo/redis"
	"github.com/honorjoey/gin-xorm/log"
	"os"
	"time"
)

var r redis.Conn

type RLog struct {
	log.Log
}

var rLog RLog

func init() {
	rLog.Tag = "Redis"

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		rLog.Error(err)
	}

	host := cfg.Section("redis").Key("host").Value()
	port := cfg.Section("redis").Key("port").Value()
	password := cfg.Section("redis").Key("password").Value()

	r, err = redis.Dial("tcp", host+":"+port, redis.DialKeepAlive(1*time.Second),
		redis.DialPassword(password),
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second))
	if err != nil {
		rLog.Error("Start redis client err", err)
		os.Exit(-1)
	}
}
