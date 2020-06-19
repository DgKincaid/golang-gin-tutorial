package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookBson used for the mongo db
type BookBson struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author string             `json:"author" bson:"author,omitempty"`
}

//FindAll books
func FindAll() ([]BookBson, error) {
	var books []BookBson

	// models.DB.Find(&books)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := MongoDB.Collection("books").Find(ctx, bson.M{})

	if err != nil {
		return books, err
	}

	if err = cursor.All(ctx, &books); err != nil {
		return books, err
	}

	return books, nil
}

// CreateBook creates book
// Create new book
func CreateBook(title string, author string) (primitive.ObjectID, error) {

	book := BookBson{Title: title, Author: author}

	// models.DB.Find(&books)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := MongoDB.Collection("books").InsertOne(ctx, book)

	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, errors.New("Unable to insert book")
}
