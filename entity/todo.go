package entity

import "gorm.io/gorm"

type Todos struct {
	Todos []TODO `json:"todos"`
}

type TODO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}

type Todo struct {
	gorm.Model
	Title      string
	TodoListID uint
	UserID     uint
}
