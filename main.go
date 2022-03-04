package main

import (
	"go-gin-todo/controller"
	"go-gin-todo/service"

	"github.com/gin-gonic/gin"
)

func main() {
	service.Init()

	var Router *gin.Engine

	Router = gin.New()
	Router.Use(gin.Logger())
	api := Router.Group("/api")
	{
		controller.BindTodoRoutes(api)
	}
	Router.Run(":3009")
}
