package controller

import (
	"go-gin-todo/entity"
	"go-gin-todo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// @Summary Get all todos
// @Accept  json
// @Produce  json
// @Success 200 {object}
// @Router /todo [get]
func GetAllTodos(ctx *gin.Context) {
	todos := service.GetAllTodos()

	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
	})

}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /todo [get]
func AddTodo(ctx *gin.Context) {
	todo := entity.TODO{}

	if err := ctx.BindJSON(&todo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTodo := service.CreateTodo(todo.Title)
	ctx.JSON(http.StatusAccepted, &newTodo)
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /todo [get]
func GetSingleTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := service.GetTodoById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusAccepted, &todo)
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /todo [get]
func UpdateTodo(ctx *gin.Context) {
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
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /todo [get]
func CompleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := service.CompleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusAccepted, &todo)
}

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /todo [get]
func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := service.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusAccepted, "")
}
