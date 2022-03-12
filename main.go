package main

import (
	"go-gin-todo/controller"
	"go-gin-todo/docs"
	"go-gin-todo/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router *gin.Engine

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

// @host localhost:3009

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	entity.ConnectDB()

	Router = gin.New()
	docs.SwaggerInfo.BasePath = "/api"
	Router.Use(gin.Logger())

	Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong")
	})

	api := Router.Group("/api")
	{
		todoRouter := api.Group("/todo")
		{
			todoRouter.GET("/", controller.GetAllTodos)
			todoRouter.POST("/", controller.AddTodo)
			todoRouter.GET("/:id", controller.GetSingleTodo)
			todoRouter.PUT("/:id", controller.UpdateTodo)
			todoRouter.PUT("/:id/complete", controller.CompleteTodo)
			todoRouter.DELETE("/:id", controller.DeleteTodo)
		}
	}
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	Router.Run(":3009")
}
