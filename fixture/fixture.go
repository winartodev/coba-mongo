package fixture

import (
	"winartodev/coba-mongodb/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Categories = []entity.Category{
		{
			ID:   primitive.NewObjectID(),
			Slug: "category-1",
			Name: "Category 1",
			SubCategory: []entity.SubCategory{
				{
					ID:   primitive.NewObjectID(),
					Slug: "sub-category-1",
					Name: "Sub Category 1",
				},
			},
		},
		{
			ID:   primitive.NewObjectID(),
			Slug: "category-2",
			Name: "Category 2",
			SubCategory: []entity.SubCategory{
				{
					ID:   primitive.NewObjectID(),
					Slug: "sub-category-2",
					Name: "Sub Category 2",
				},
			},
		},
		{
			ID:   primitive.NewObjectID(),
			Slug: "category-3",
			Name: "Category 3",
			SubCategory: []entity.SubCategory{
				{
					ID:   primitive.NewObjectID(),
					Slug: "sub-category-3",
					Name: "Sub Category 3",
				},
			},
		},
	}

	Category = entity.Category{
		ID:   primitive.NewObjectID(),
		Slug: "category-1",
		Name: "Category 1",
		SubCategory: []entity.SubCategory{
			{
				ID:   primitive.NewObjectID(),
				Slug: "sub-category-1",
				Name: "Sub Category 1",
			},
		},
	}

	CategoryWithNillObjectID = entity.Category{
		ID:   primitive.NilObjectID,
		Slug: "category-1",
		Name: "Category",
		SubCategory: []entity.SubCategory{
			{
				ID:   primitive.NilObjectID,
				Slug: "sub-category-1",
				Name: "Sub Category 1",
			},
		},
	}
)
