package entity

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DB_DSN")

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection.AutoMigrate(&Todo{})
	connection.AutoMigrate(&TodoList{})
	connection.AutoMigrate(&User{})

	DB = connection

}
