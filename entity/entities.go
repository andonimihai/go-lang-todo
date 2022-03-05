package entity

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=postgres password=postgres dbname=todo port=5432 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection.AutoMigrate(&Todo{})
	connection.AutoMigrate(&TodoList{})
	connection.AutoMigrate(&User{})

	DB = connection

}
