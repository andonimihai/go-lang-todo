package entity

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	Title  string
	UserID uint
	Todos  []Todo `gorm:"foreignKey:TodoListID"`
}
