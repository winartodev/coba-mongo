package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShippingAddress struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Order struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	UserID          primitive.ObjectID `json:"user_id" bson:"user_id"`
	State           string             `json:"state"`
	LineItems       []Product          `json:"line_items" bson:"line_items"`
	ShippingAddress ShippingAddress    `json:"shipping_address" bson:"shipping_address"`
	SubTotal        int                `json:"sub_total"`
}
