package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/honorjoey/gin-xorm/middleware"
	"net/http"
)

type MsgController struct {
	middleware.MsgSender
}

type MsgData struct {
	Topic string `json:"topic"`
	Msg string `json:"msg"`
}

// @Summary Send kafka message
// @Description kafka producer send message
// @Accept  json
// @Produce  json
// @Param topic query string true "topic"
// @Param msg query string true "msg"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /api/send [post]
func (m MsgController) SendMessage(ctx *gin.Context)  {
	data := new(MsgData)
	ctx.Bind(data)
	err := m.SendMsg(data.Topic, data.Msg)
	if err != nil {
		cLog.Error("send message err", err)
	}
	ctx.SecureJSON(http.StatusOK, gin.H{"msg":"ok"})
}
