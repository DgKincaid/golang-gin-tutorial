package models

import (
	"github.com/jinzhu/gorm"

	// comment justifying this import
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB test database setup
var DB *gorm.DB

// ConnectDB connect to test db
func ConnectDB() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
