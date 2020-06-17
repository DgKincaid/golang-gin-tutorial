package models

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"

	// comment justifying this import
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// MongoDB instance of the mongo database
var MongoDB *mongo.Database

// ConnectMongoDB connects to a mongo db
func ConnectMongoDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:ZY+YCzCMlMUB0TdALcNLSQQbGGSCV7hA=@localhost:27017/test"))

	if err != nil {
		panic("Failed to connect to mongo DB")
	}

	defer client.Disconnect(ctx)

	database := client.Database("test")
	database.Collection("books")

	MongoDB = database
}
