package service

import (
	"go-gin-todo/entity"
)

type TodoState string

const (
	Open      TodoState = "open"
	Completed TodoState = "completed"
)

func GetAllTodos() []entity.Todo {
	var todos []entity.Todo
	entity.DB.Find(&todos).Limit(10).Offset(0)
	return todos
}

func CreateTodo(title string) entity.Todo {

	todo := entity.Todo{Title: title, State: string(Open)}

	result := entity.DB.Create(&todo)

	if result.Error != nil {
		panic(result.Error)
	}

	return todo
}

func GetTodoById(id string) (entity.Todo, error) {
	var todo entity.Todo
	result := entity.DB.First(&todo, id)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil

}

func UpdateTodoName(id string, newTitle string) (entity.Todo, error) {
	var todo entity.Todo
	result := entity.DB.First(&todo, id)

	if result.Error != nil {
		return todo, result.Error
	}

	todo.Title = newTitle
	saveResult := entity.DB.Save(&todo)

	if saveResult.Error != nil {
		return todo, saveResult.Error
	}

	return todo, nil
}

func CompleteTodo(id string) (entity.Todo, error) {
	var todo entity.Todo
	result := entity.DB.First(&todo, id)

	if result.Error != nil {
		return todo, result.Error
	}

	todo.State = string(Completed)
	saveResult := entity.DB.Save(&todo)

	if saveResult.Error != nil {
		return todo, saveResult.Error
	}

	return todo, nil

}

func DeleteTodo(id string) error {
	result := entity.DB.Delete(&entity.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
