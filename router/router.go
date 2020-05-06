package router

import (
	ginsessions "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/honorjoey/gin-xorm/controller"
	_ "github.com/honorjoey/gin-xorm/docs"
	"github.com/honorjoey/gin-xorm/router/middlewares"
	"github.com/honorjoey/gin-xorm/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() *gin.Engine {
	router := gin.New()

	cookieStore := sessions.NewCookieStore([]byte("MTU4ODUxNTQ4NnxEdi1CQkFFQ180SUFBUkFCRUFBQUJQLUNBQUE9fHRkSpeeNYnec4d2dJA4hPIsXgTJbyBDljFRehZOVS4w"))
	store := utils.NewStore(cookieStore)
	router.Use(ginsessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandler)

	router.NoRoute(middlewares.NoRoute)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := router.Group("/api")
	{
		user := new(controller.UserController)
		group.GET("/user", user.List)

		name := new(controller.NameController)
		group.POST("/name", name.GetName)
	}

	return router
}
