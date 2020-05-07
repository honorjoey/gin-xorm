package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/honorjoey/gin-xorm/database"
	"net/http"
)

type UserController struct {
	UserDao database.UserDao
}

// @Summary List users
// @Description list user
// @Produce  json
// @Param phone query string true "phone"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /api/user [get]
func (c UserController) List(ctx *gin.Context) {
	phone := ctx.Query("phone")
	user := *c.UserDao.QueryUser(phone)

	ctx.SecureJSON(http.StatusOK, user[0])
}
