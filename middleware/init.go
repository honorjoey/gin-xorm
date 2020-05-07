package middleware

import (
	"github.com/go-ini/ini"
	"github.com/honorjoey/gin-xorm/log"
)

type KLog struct {
	log.Log
}

var kLog KLog

var host, port string

func init() {
	kLog.Tag = "kafka"

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		kLog.Error(err)
	}

	host = cfg.Section("kafka").Key("host").Value()
	port = cfg.Section("kafka").Key("port").Value()

	StartProducer()
	StartConsumer()

	rev := new(MsgReceiver)
	rev.GetMsg("web")
}
