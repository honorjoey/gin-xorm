package router

import (
	ginSessions "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/honorjoey/gin-xorm/controller"
	_ "github.com/honorjoey/gin-xorm/docs"
	"github.com/honorjoey/gin-xorm/router/ginmiddleware"
	"github.com/honorjoey/gin-xorm/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() *gin.Engine {
	router := gin.New()

	cookieStore := sessions.NewCookieStore([]byte("ss"))
	store := utils.NewStore(cookieStore)
	router.Use(ginSessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(ginmiddleware.ErrorHandler)

	//router.NoRoute(ginmiddleware.NoRoute)

	router.Use(ginmiddleware.RequestTime)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := router.Group("/api")
	{
		user := new(controller.UserController)
		group.GET("/user", user.List)

		name := new(controller.NameController)
		group.POST("/name", name.GetName)

		sender := new(controller.MsgController)
		group.POST("/send", sender.SendMessage)
	}

	return router
}
