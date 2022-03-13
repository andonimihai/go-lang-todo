package controller

import (
	"go-gin-todo/entity"
	"go-gin-todo/helper"
	validator "go-gin-todo/lib"
	"go-gin-todo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Get all todos
// @Description Fetch a list of Todos that belongs to logged in user
// @Produce  json
// @Success 200 {array} entity.Todo
// @Router /todo [get]
func GetAllTodos(ctx *gin.Context) {
	user := helper.GetLoggedInUser(ctx)
	todos := service.GetAllTodos(user)

	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
	})

}

// @Security ApiKeyAuth
// @Summary Add a new todo
// @Description Add a new todo to the list
// @Accept  json
// @Produce  json
// @Param todo body entity.UpsertTodo true "Add new todo"
// @Success 201 {object} entity.Todo
// @Router /todo [post]
func AddTodo(ctx *gin.Context) {
	user := helper.GetLoggedInUser(ctx)
	var todo entity.UpsertTodo
	if err := ctx.ShouldBind(&todo); err != nil {
		validator.HandleBindingError(err, ctx)
		return
	}
	newTodo := service.CreateTodo(todo.Title, user)
	ctx.JSON(http.StatusCreated, &newTodo)
}

// @Security ApiKeyAuth
// @Summary Fetch single Todo Item
// @Description get Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Success 200 {object} entity.Todo
// @Router /todo/{id} [get]
func GetSingleTodo(ctx *gin.Context) {
	user := helper.GetLoggedInUser(ctx)
	id := ctx.Param("id")
	todo, err := service.GetTodoById(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &todo)
}

// @Security ApiKeyAuth
// @Summary Update existing Todo
// @Description update Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Param todo body entity.UpsertTodo true "Update todo"
// @Success 200 {object} entity.Todo
// @Router /todo/{id} [put]
func UpdateTodo(ctx *gin.Context) {
	user := helper.GetLoggedInUser(ctx)
	id := ctx.Param("id")
	todo := entity.UpsertTodo{}

	if err := ctx.ShouldBind(&todo); err != nil {
		validator.HandleBindingError(err, ctx)
		return
	}

	updatedTodo, err := service.UpdateTodoName(id, todo.Title, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &updatedTodo)
}

// @Security ApiKeyAuth
// @Summary Complete a Todo
// @Description complete Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Success 200 {object} entity.Todo
// @Router /todo/{id}/complete [put]
func CompleteTodo(ctx *gin.Context) {
	user := helper.GetLoggedInUser(ctx)
	id := ctx.Param("id")
	todo, err := service.CompleteTodo(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &todo)
}

// @Security ApiKeyAuth
// @Summary Delete Todo
// @Description delete Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Success 200 {string} string	"ok"
// @Router /todo/{id} [delete]
func DeleteTodo(ctx *gin.Context) {
	user := helper.GetLoggedInUser(ctx)
	id := ctx.Param("id")
	err := service.DeleteTodo(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}
