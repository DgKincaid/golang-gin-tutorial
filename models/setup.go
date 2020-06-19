package models

import (
	"context"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	// comment justifying this import
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	clientOptions := options.Client().ApplyURI("mongodb://admin:pwd123@127.0.0.1:27017/gomongo")
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Failed to connect to mongo DB", err)
	} else {
		log.Println("Database connected")
	}

	database := client.Database("gomongo")

	MongoDB = database
}
