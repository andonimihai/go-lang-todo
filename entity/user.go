package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	EmailName string
	Todos     []Todo     `gorm:"foreignKey:UserID"`
	TodoLists []TodoList `gorm:"foreignKey:UserID"`
}
