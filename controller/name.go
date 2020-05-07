package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/honorjoey/gin-xorm/database"
	"net/http"
)

type NameController struct {
	NameDao database.NameDao
}

type NameData struct {
	Data string `json:"data"`
}

// @Summary Get name value
// @Description get name value
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /api/name [post]
func (n NameController) GetName(ctx *gin.Context) {
	name := NameData{}
	ctx.Bind(&name)

	resp := n.NameDao.GetName(name.Data)

	ctx.SecureJSON(http.StatusOK, resp)
}
