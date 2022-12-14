package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("sqlite3", "app.db")

	if err != nil {
		panic("Failed to connect to the Database!")

	}
	db.AutoMigrate(&User{})

	DB = db
}
