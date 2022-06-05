package repository_test

import (
	"context"
	"errors"
	"testing"
	"winartodev/coba-mongodb/entity"
	"winartodev/coba-mongodb/fixture"
	"winartodev/coba-mongodb/repository"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFindAll(t *testing.T) {
	testcases := []struct {
		name      string
		data      []entity.Category
		wantError bool
		err       error
	}{
		{
			name: "success",
			data: fixture.Categories,
		},
		{
			name:      "failed when collect data",
			wantError: true,
			err:       errors.New("failed when collect data"),
		},
	}
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, test := range testcases {
		mt.Run(test.name, func(mt *mtest.T) {
			categoryRepo := repository.NewCategoryRepository(mt.DB)

			if !test.wantError {
				data := mtest.CreateCursorResponse(2, "products.categories", mtest.FirstBatch, bson.D{
					primitive.E{Key: "_id", Value: test.data[0].ID},
					primitive.E{Key: "slug", Value: test.data[0].Slug},
					primitive.E{Key: "name", Value: test.data[0].Name},
					primitive.E{Key: "sub_category", Value: test.data[0].SubCategory},
				})

				killCursors := mtest.CreateCursorResponse(0, "products.categories", mtest.NextBatch)
				mt.AddMockResponses(data, killCursors)
			}

			res, err := categoryRepo.FindAll(context.Background())
			if err != nil && test.wantError {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestFindOne(t *testing.T) {
	testcases := []struct {
		name      string
		slug      string
		data      entity.Category
		wantError bool
		err       error
	}{
		{
			name: "success",
			slug: "category-1",
			data: fixture.Category,
		},
		{
			name:      "failed",
			slug:      "category-2",
			data:      entity.Category{},
			wantError: true,
			err:       mongo.ErrNoDocuments,
		},
	}
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, test := range testcases {
		mt.Run(test.name, func(mt *mtest.T) {
			categoryRepo := repository.NewCategoryRepository(mt.DB)
			if !test.wantError {
				data := mtest.CreateCursorResponse(1, "products.categories", mtest.FirstBatch, bson.D{
					primitive.E{Key: "_id", Value: test.data.ID},
					primitive.E{Key: "slug", Value: test.data.Slug},
					primitive.E{Key: "name", Value: test.data.Name},
					primitive.E{Key: "sub_category", Value: test.data.SubCategory},
				})
				killCursors := mtest.CreateCursorResponse(0, "products.categories", mtest.NextBatch)
				mt.AddMockResponses(data, killCursors)
			} else {
				mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{Message: test.err.Error()}))
			}

			res, err := categoryRepo.FindOne(context.Background(), test.slug)
			if err != nil && test.wantError {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	testcases := []struct {
		name      string
		data      entity.Category
		wantError bool
		err       error
	}{
		{
			name: "success",
			data: fixture.Category,
		},
		{
			name:      "failed",
			wantError: true,
			err:       errors.New("failed to insert data"),
		},
	}
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, test := range testcases {
		mt.Run(test.name, func(mt *mtest.T) {
			categoryRepository := repository.NewCategoryRepository(mt.DB)

			if !test.wantError {
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			} else {
				mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
					Index:   1,
					Code:    200,
					Message: test.err.Error(),
				}))
			}

			err := categoryRepository.Insert(context.Background(), test.data)
			if err != nil && test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	testcases := []struct {
		name      string
		slug      string
		data      entity.Category
		wantError bool
		err       error
	}{
		{
			name: "success",
			slug: "category-1",
			data: fixture.CategoryWithNillObjectID,
		},
		{
			name:      "failed",
			wantError: true,
			err:       errors.New("failed when update data"),
		},
	}
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, test := range testcases {
		mt.Run(test.name, func(mt *mtest.T) {
			categoryRepository := repository.NewCategoryRepository(mt.DB)

			if !test.wantError {
				mt.AddMockResponses(bson.D{
					primitive.E{Key: "ok", Value: 1},
					primitive.E{Key: "data", Value: bson.D{
						primitive.E{Key: "_id", Value: test.data.ID},
						primitive.E{Key: "slug", Value: test.data.Slug},
						primitive.E{Key: "name", Value: test.data.Name},
						primitive.E{Key: "sub_category", Value: test.data.SubCategory},
					}},
				})
			} else {
				mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
					Index:   1,
					Code:    200,
					Message: test.err.Error(),
				}))
			}

			err := categoryRepository.Update(context.Background(), test.slug, test.data)
			if err != nil && test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testcases := []struct {
		name      string
		slug      string
		wantError bool
		err       error
	}{
		{
			name: "success",
			slug: "category-1",
		},
		{
			name:      "failed",
			slug:      "category-1",
			wantError: true,
			err:       errors.New("failed when delete data"),
		},
	}

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	for _, test := range testcases {
		mt.Run(test.name, func(mt *mtest.T) {
			categoryRepository := repository.NewCategoryRepository(mt.DB)

			if !test.wantError {
				mt.AddMockResponses(bson.D{
					primitive.E{Key: "ok", Value: 1},
				})
			} else {
				mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
					Index:   1,
					Code:    200,
					Message: test.err.Error(),
				}))
			}

			err := categoryRepository.Delete(context.Background(), test.slug)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
