package router

import (
	"github.com/honorjoey/gin-xorm/controller"
	"github.com/honorjoey/gin-xorm/router/middlewares"
	"github.com/honorjoey/gin-xorm/utils"
	ginsessions "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func Init() *gin.Engine {
	router := gin.New()

	cookieStore := sessions.NewCookieStore([]byte("jDIkFg6ju7kEM7DOIWGcXSLwCL6QaMZy"))
	store := utils.NewStore(cookieStore)
	router.Use(ginsessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandler)

	router.NoRoute(middlewares.NoRoute)

	user := new(controller.UserController)
	router.GET("/user", user.List)

	return router
}
