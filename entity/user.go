package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
}

type PaymentMethod struct {
}

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Username       string             `json:"username"`
	Email          string             `json:"email"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	HashedPassword string             `json:"hashed_password"`
	Address        []Address          `json:"address" bson:"address"`
	PaymentMethod  []PaymentMethod    `json:"payment_methods" bson:"payment_methods"`
}
