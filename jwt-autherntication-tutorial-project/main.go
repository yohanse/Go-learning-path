package main

import (
	"example/jwt-authentication-tutorial-project/controller"
	"example/jwt-authentication-tutorial-project/database"
	"example/jwt-authentication-tutorial-project/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("postgresql://postgres:postgres@localhost:5432/jwtautherntication")
	router := initRouter()
  	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
	  api.POST("/token", controller.GenerateToken)
	  api.POST("/user/register", controller.RegisterUser)
	  secured := api.Group("/secured").Use(middleware.Auth())
	  {
		secured.GET("/ping", controller.Ping)
	  }
	}
	return router
  }