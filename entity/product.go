package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Detail struct {
	Weight       int    `json:"weight"`
	WeightUnits  string `json:"weight_units"`
	ModelNum     int    `json:"model_num"`
	Manufacturer string `json:"manufacturer"`
	Color        string `json:"color"`
}

type Pricing struct {
	Retail int `json:"retail"`
	Sale   int `json:"sale"`
}

type PriceHistory struct {
	Retail int       `json:"retail"`
	Sale   int       `json:"sale"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}

type Product struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Slug            string             `json:"slug"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Detail          Detail             `json:"details" bson:"details"`
	TotalReviews    int                `json:"total_reviews"`
	AverageReviews  float32            `json:"average_review"`
	Pricing         Pricing            `json:"pricing" bson:"pricing"`
	PriceHistory    []PriceHistory     `json:"price_history" bson:"price_history"`
	PrimaryCategory primitive.ObjectID `json:"primary_category" bson:"primary_category"`
	Tags            []string           `json:"tags"`
}
