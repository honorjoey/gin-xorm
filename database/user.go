package database

import (
	"github.com/honorjoey/gin-xorm/models"
)

// http://www.admpub.com:8080/xorm-manual-zh-CN/

type UserDao struct {}

func (UserDao) QueryUser(phone string) *[]models.User {
	var user []models.User
	err := x.Where("telphone=?", phone).Find(&user)
	if err != nil {
		xLog.Error("QueryUser ", err)
	}
	xLog.Info("Successful")
	return &user
}
