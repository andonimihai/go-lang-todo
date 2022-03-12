package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title      string
	State      string
	TodoListID uint
	UserID     uint
}

type UpsertTodo struct {
	Title string `json:"title" binding:"required"`
}
