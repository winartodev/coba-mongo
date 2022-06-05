package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubCategory struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Slug string             `json:"slug"`
	Name string             `json:"name"`
}

type Category struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Slug        string             `json:"slug"`
	Name        string             `json:"name"`
	SubCategory []SubCategory      `json:"sub_category" bson:"sub_category"`
}
