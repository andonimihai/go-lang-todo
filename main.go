package main

import (
	"go-gin-todo/controller"
	"go-gin-todo/docs"
	"go-gin-todo/entity"
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

// @title Simple Todo API
// @version 1.0
// @description This is a server to manage todos
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
