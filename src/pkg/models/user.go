package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// ID Generates Automatically - MongoDB Save ID as _id
	// ID ~> Golang Server, id ~> API (POSTMAN), _id ~> MongoDB
	ID     primitive.ObjectID `json:"id"     bson:"_id"`
	Name   string             `json:"name"   bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age"    bson:"age"`
}
