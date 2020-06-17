package models

// BookBson used for the mongo db
type BookBson struct {
	ID     uint `bson:"_id,omitempty"`
	Title  uint `bson:"title,omitempty"`
	Author uint `bson:"author,omitempty"`
}
