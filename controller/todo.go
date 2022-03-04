package controller

import (
	"go-gin-todo/entity"
	"go-gin-todo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindTodoRoutes(route *gin.RouterGroup) {
	todoRouter := route.Group("/todo")
	{
		todoRouter.GET("/", func(ctx *gin.Context) {
			todos := service.GetAllTodos()
			ctx.JSON(http.StatusOK, gin.H{
				"data": todos.Todos,
			})
		})

		todoRouter.POST("/", func(ctx *gin.Context) {
			todo := entity.TODO{}

			if err := ctx.BindJSON(&todo); err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			newTodo := service.CreateTodo(todo.Title)
			ctx.JSON(http.StatusAccepted, &newTodo)
		})

		todoRouter.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			todo, err := service.GetTodoById(id)
			if err != nil {
				ctx.JSON(http.StatusNotFound, err)
				return
			}

			ctx.JSON(http.StatusAccepted, &todo)
		})

		todoRouter.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			todo := entity.TODO{}

			if err := ctx.BindJSON(&todo); err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			updatedTodo, err := service.UpdateTodoName(id, todo.Title)
			if err != nil {
				ctx.JSON(http.StatusNotFound, err)
				return
			}

			ctx.JSON(http.StatusAccepted, &updatedTodo)
		})

		todoRouter.PUT("/:id/complete", func(ctx *gin.Context) {
			id := ctx.Param("id")
			todo, err := service.CompleteTodo(id)
			if err != nil {
				ctx.JSON(http.StatusNotFound, err)
				return
			}

			ctx.JSON(http.StatusAccepted, &todo)
		})

		todoRouter.DELETE("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			_, err := service.DeleteTodo(id)
			if err != nil {
				ctx.JSON(http.StatusNotFound, err)
				return
			}

			ctx.JSON(http.StatusAccepted, "")
		})
	}
}
