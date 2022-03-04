package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todos struct {
	Todos []TODO `json:"todos"`
}

type TODO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("todo.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened todo.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var todos Todos

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &todos)

	var Router *gin.Engine

	Router = gin.Default()
	api := Router.Group("/api")
	{
		api.GET("/todo", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"todos": todos.Todos,
			})
		})

		api.POST("/todo", func(ctx *gin.Context) {
			todo := TODO{}
			id := uuid.New()
			if err := ctx.BindJSON(&todo); err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
			todo.ID = id.String()
			todo.State = "open"
			todos.Todos = append(todos.Todos, todo)

			saveTodos(&todos)

			ctx.JSON(http.StatusAccepted, &todo)
		})

		api.GET("/todo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			var foundTodo TODO
			var found = false

			for _, todo := range todos.Todos {
				if todo.ID == id {
					foundTodo = todo
					found = true
					break
				}
			}

			if found {
				ctx.JSON(http.StatusOK, &foundTodo)
			} else {
				ctx.JSON(http.StatusNotFound, "")
			}

		})

		api.PUT("/todo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			var todoIdx = -1
			todo := TODO{}

			if err := ctx.BindJSON(&todo); err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			for idx, todo := range todos.Todos {
				fmt.Println(idx)
				if todo.ID == id {
					todoIdx = idx
					break
				}
			}

			if todoIdx == -1 {
				ctx.JSON(http.StatusNotFound, "")
				return
			}
			todos.Todos[todoIdx].Title = todo.Title

			saveTodos(&todos)

			ctx.JSON(http.StatusAccepted, &todos.Todos[todoIdx])

		})

		api.PUT("/todo/:id/complete", func(ctx *gin.Context) {
			id := ctx.Param("id")
			var todoIdx = -1

			for idx, todo := range todos.Todos {
				fmt.Println(idx)
				if todo.ID == id {
					todoIdx = idx
					break
				}
			}

			if todoIdx == -1 {
				ctx.JSON(http.StatusNotFound, "")
				return
			}
			todos.Todos[todoIdx].State = "completed"

			saveTodos(&todos)

			ctx.JSON(http.StatusAccepted, &todos.Todos[todoIdx])
		})

		api.DELETE("/todo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			var todoIdx = -1

			for idx, todo := range todos.Todos {
				fmt.Println(idx)
				if todo.ID == id {
					todoIdx = idx
					break
				}
			}

			if todoIdx == -1 {
				ctx.JSON(http.StatusNotFound, "")
				return
			}
			todos.Todos[todoIdx].State = "deleted"

			saveTodos(&todos)

			ctx.JSON(http.StatusAccepted, &todos.Todos[todoIdx])
		})
	}
	Router.Run(":3009")
}

func saveTodos(todos *Todos) {
	file, _ := json.MarshalIndent(todos, "", " ")

	_ = ioutil.WriteFile("todo.json", file, 0644)
}
