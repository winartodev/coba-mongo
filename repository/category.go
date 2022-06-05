package repository

import (
	"context"
	"errors"
	"winartodev/coba-mongodb/entity"

	s "github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CategoryCollectionName = "categories"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]entity.Category, error)
	FindOne(ctx context.Context, slug string) (*entity.Category, error)
	Insert(ctx context.Context, category entity.Category) error
	Update(ctx context.Context, slug string, category entity.Category) error
	Delete(ctx context.Context, slug string) error
}

type category struct {
	db *mongo.Database
}

func NewCategoryRepository(db *mongo.Database) CategoryRepository {
	return &category{db: db}
}

func (c *category) FindAll(ctx context.Context) ([]entity.Category, error) {
	categories := []entity.Category{}

	cursor, err := c.db.Collection(CategoryCollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		category := entity.Category{}
		err := cursor.Decode(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *category) FindOne(ctx context.Context, slug string) (*entity.Category, error) {
	category := entity.Category{}

	err := c.db.Collection(CategoryCollectionName).FindOne(ctx, bson.D{primitive.E{Key: "slug", Value: slug}}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("data has not found")
		}
		return nil, err
	}
	return &category, nil
}

func (c *category) Insert(ctx context.Context, category entity.Category) error {
	category.ID = primitive.NewObjectID()
	category.Slug = s.Make(category.Name)

	for i := range category.SubCategory {
		category.SubCategory[i].ID = primitive.NewObjectID()
		category.SubCategory[i].Slug = s.Make(category.SubCategory[i].Name)
	}

	newCategory := bson.D{
		primitive.E{Key: "slug", Value: category.Slug},
		primitive.E{Key: "name", Value: category.Name},
		primitive.E{Key: "sub_category", Value: category.SubCategory},
	}

	_, err := c.db.Collection(CategoryCollectionName).InsertOne(ctx, newCategory)
	if err != nil {
		return err
	}

	return nil
}

func (c *category) Update(ctx context.Context, slug string, category entity.Category) error {
	category.ID = primitive.NewObjectID()
	category.Slug = s.Make(category.Name)

	for i := range category.SubCategory {
		if category.SubCategory[i].ID == primitive.NilObjectID {
			category.SubCategory[i].ID = primitive.NewObjectID()
		}
		category.SubCategory[i].Slug = s.Make(category.SubCategory[i].Name)
	}

	newCategory := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "slug", Value: category.Slug},
		primitive.E{Key: "name", Value: category.Name},
		primitive.E{Key: "sub_category", Value: category.SubCategory},
	}}}

	_, err := c.db.Collection(CategoryCollectionName).UpdateOne(ctx, bson.D{primitive.E{Key: "slug", Value: slug}}, newCategory)
	if err != nil {
		return err
	}

	return nil
}

func (c *category) Delete(ctx context.Context, slug string) error {
	_, err := c.db.Collection(CategoryCollectionName).DeleteOne(ctx, bson.D{primitive.E{Key: "slug", Value: slug}})
	if err != nil {
		return err
	}
	return nil
}
