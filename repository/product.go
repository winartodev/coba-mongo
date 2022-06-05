package repository

import (
	"context"
	"winartodev/coba-mongodb/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ProductCollectionName = "products"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	Insert(ctx context.Context, product entity.Product) error
}

type productRepository struct {
	db *mongo.Database
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	products := []entity.Product{}
	curr, err := r.db.Collection(ProductCollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for curr.Next(ctx) {
		product := entity.Product{}

		err := curr.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) Insert(ctx context.Context, product entity.Product) error {
	product.ID = primitive.NewObjectID()

	_, err := r.db.Collection(ProductCollectionName).InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
