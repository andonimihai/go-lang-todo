package controller

import (
	"go-gin-todo/entity"
	"go-gin-todo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all todos
// @Description Fetch a list of Todos that belongs to logged in user
// @Produce  json
// @Success 200 {array} entity.Todo
// @Router /todo [get]
func GetAllTodos(ctx *gin.Context) {
	todos := service.GetAllTodos()

	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
	})

}

// @Summary Add a new todo
// @Description Add a new todo to the list
// @Accept  json
// @Produce  json
// @Param todo body entity.UpsertTodo true "Add new todo"
// @Success 201 {object} entity.Todo
// @Router /todo [post]
func AddTodo(ctx *gin.Context) {
	todo := entity.TODO{}

	if err := ctx.BindJSON(&todo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTodo := service.CreateTodo(todo.Title)
	ctx.JSON(http.StatusCreated, &newTodo)
}

// @Summary Fetch single Todo Item
// @Description get Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Success 200 {object} entity.Todo
// @Router /todo/{id} [get]
func GetSingleTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := service.GetTodoById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &todo)
}

// @Summary Update existing Todo
// @Description update Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Param todo body entity.UpsertTodo true "Update todo"
// @Success 200 {object} entity.Todo
// @Router /todo/{id} [put]
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

	ctx.JSON(http.StatusOK, &updatedTodo)
}

// @Summary Complete a Todo
// @Description complete Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Success 200 {object} entity.Todo
// @Router /todo/{id}/complete [put]
func CompleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := service.CompleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &todo)
}

// @Summary Delete Todo
// @Description delete Todo by ID
// @Accept  json
// @Produce  json
// @Param id path int  true "Todo ID"
// @Success 200 {string} string	"ok"
// @Router /todo/{id} [delete]
func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := service.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, "")
}
