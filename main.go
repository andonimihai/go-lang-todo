package main

import (
	"go-gin-todo/controller"
	"go-gin-todo/docs"
	"go-gin-todo/entity"
	"go-gin-todo/lib"
	"go-gin-todo/middleware"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

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
	lib.InitFirebaseAuth()

	Router = gin.New()
	docs.SwaggerInfo.BasePath = "/api"

	Router.Use(gin.Logger())
	// using the auth middle ware to validate api requests
	Router.Use(middleware.AuthMiddleware)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

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
