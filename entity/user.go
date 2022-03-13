package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"uniqueIndex"`
	ImageUrl   string
	ExternalId string
	Todos      []Todo     `gorm:"foreignKey:UserID"`
	TodoLists  []TodoList `gorm:"foreignKey:UserID"`
}

type UpsertUser struct {
	Name   string
	UserId string
	Email  string
	Avatar string
}
