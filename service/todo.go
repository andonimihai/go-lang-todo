package service

import (
	"go-gin-todo/entity"
)

type TodoState string

const (
	Open      TodoState = "open"
	Completed TodoState = "completed"
)

func GetAllTodos(user entity.User) []entity.Todo {
	var todos []entity.Todo
	entity.DB.Where(&entity.Todo{UserID: user.ID}).Find(&todos).Limit(10).Offset(0)
	return todos
}

func CreateTodo(title string, user entity.User) entity.Todo {

	todo := entity.Todo{Title: title, State: string(Open), UserID: user.ID}

	result := entity.DB.Create(&todo)

	if result.Error != nil {
		panic(result.Error)
	}

	return todo
}

func GetTodoById(id string, user entity.User) (entity.Todo, error) {
	var todo entity.Todo
	result := entity.DB.Where(&entity.Todo{UserID: user.ID}).First(&todo, id)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil

}

func UpdateTodoName(id string, newTitle string, user entity.User) (entity.Todo, error) {
	var todo entity.Todo
	result := entity.DB.Where(&entity.Todo{UserID: user.ID}).First(&todo, id)

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

func CompleteTodo(id string, user entity.User) (entity.Todo, error) {
	var todo entity.Todo
	result := entity.DB.Where(&entity.Todo{UserID: user.ID}).First(&todo, id)

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

func DeleteTodo(id string, user entity.User) error {
	result := entity.DB.Where("user_id = ?", user.ID).Where("id = ?", id).Delete(&entity.Todo{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
