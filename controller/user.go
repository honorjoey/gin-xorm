package controller

import (
	"github.com/honorjoey/gin-xorm/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserDao db.UserDao
}

func (c UserController) List(ctx *gin.Context) {
	user := c.UserDao.QueryUser()

	ctx.SecureJSON(http.StatusOK, user)
}
