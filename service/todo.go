package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-gin-todo/entity"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

var todoFileName = "todo.json"

var todos entity.Todos

func Init() {
	// Open our jsonFile
	jsonFile, err := os.Open(todoFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened todo.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'todos' which we defined above
	json.Unmarshal(byteValue, &todos)
}

func GetAllTodos() []entity.Todo {
	var todos []entity.Todo
	entity.DB.Find(&todos).Limit(10).Offset(0)
	return todos
}

func saveTodos(todos *entity.Todos) {
	file, _ := json.MarshalIndent(todos, "", " ")

	_ = ioutil.WriteFile(todoFileName, file, 0644)

}

func CreateTodo(title string) entity.TODO {
	id := uuid.New()
	var todo entity.TODO
	todo.ID = id.String()
	todo.State = "open"
	todo.Title = title
	todos.Todos = append(todos.Todos, todo)

	saveTodos(&todos)
	return todo
}

func GetTodoById(id string) (entity.TODO, error) {
	var foundTodo entity.TODO
	found := false
	for _, todo := range todos.Todos {
		if todo.ID == id {
			foundTodo = todo
			found = true
			break
		}
	}

	if found {
		return foundTodo, nil
	}

	return foundTodo, errors.New("not Found")

}

func UpdateTodoName(id string, newTitle string) (entity.TODO, error) {
	todoIdx := -1

	for idx, todo := range todos.Todos {
		fmt.Println(idx)
		if todo.ID == id {
			todoIdx = idx
			break
		}
	}

	if todoIdx == -1 {
		return entity.TODO{}, errors.New("not found")
	}

	todos.Todos[todoIdx].Title = newTitle

	saveTodos(&todos)

	return todos.Todos[todoIdx], nil

}

func CompleteTodo(id string) (entity.TODO, error) {

	todoIdx := -1

	for idx, todo := range todos.Todos {
		fmt.Println(idx)
		if todo.ID == id {
			todoIdx = idx
			break
		}
	}

	if todoIdx == -1 {
		return entity.TODO{}, errors.New("not found")
	}

	todos.Todos[todoIdx].State = "completed"

	saveTodos(&todos)

	return todos.Todos[todoIdx], nil

}

func DeleteTodo(id string) (entity.TODO, error) {
	todoIdx := -1

	for idx, todo := range todos.Todos {
		fmt.Println(idx)
		if todo.ID == id {
			todoIdx = idx
			break
		}
	}

	if todoIdx == -1 {
		return entity.TODO{}, errors.New("not found")
	}

	todos.Todos[todoIdx].State = "deleted"

	saveTodos(&todos)

	return todos.Todos[todoIdx], nil

}
