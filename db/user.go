package db

import (
	"github.com/honorjoey/gin-xorm/models"
)

// http://www.admpub.com:8080/xorm-manual-zh-CN/

type UserDao struct {}

func (UserDao) QueryUser() *[]models.User {
	var user []models.User
	err := x.Find(&user)
	if err != nil {
		dbLogger.Error("QueryUser ", err)
	}
	return &user
}
