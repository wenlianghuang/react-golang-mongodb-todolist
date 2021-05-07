package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// "omitempty" is omitted the value if it is empty in the json or bson
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}
