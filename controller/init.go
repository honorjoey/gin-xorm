package controller

import "github.com/honorjoey/gin-xorm/log"

type CLog struct {
	log.Log
}

var cLog CLog

func init() {
	cLog.Tag = "controller"
}
