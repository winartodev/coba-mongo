package usecase_test

import (
	"context"
	"errors"
	"testing"
	"winartodev/coba-mongodb/entity"
	"winartodev/coba-mongodb/fixture"
	"winartodev/coba-mongodb/mocks"
	"winartodev/coba-mongodb/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
			name:      "failed",
			data:      nil,
			wantError: true,
			err:       errors.New("fail when get all data"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockCategoryRepo := mocks.NewCategoryRepository(t)
			mockCategoryRepo.On("FindAll", mock.Anything).Return(test.data, test.err)

			categoryUscase := usecase.NewCategoryUsecase(mockCategoryRepo)

			res, err := categoryUscase.FindAll(context.Background())
			if err != nil && test.wantError {
				assert.Nil(t, res)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindOne(t *testing.T) {
	testcases := []struct {
		name      string
		slug      string
		data      *entity.Category
		wantError bool
		err       error
	}{
		{
			name: "success",
			slug: "category-1",
			data: &fixture.Category,
		},
		{
			name:      "failed",
			slug:      "category-1",
			data:      nil,
			wantError: true,
			err:       errors.New("fail when get data"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockCategoryRepository := mocks.NewCategoryRepository(t)
			mockCategoryRepository.On("FindOne", mock.Anything, mock.AnythingOfType("string")).Return(test.data, test.err)

			categoryUsecase := usecase.NewCategoryUsecase(mockCategoryRepository)
			res, err := categoryUsecase.FindOne(context.Background(), test.slug)
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
			err:       errors.New("fail when insert data"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockCategoryRepository := mocks.NewCategoryRepository(t)
			mockCategoryRepository.On("Insert", mock.Anything, mock.AnythingOfType("entity.Category")).Return(test.err)

			categoryUsecase := usecase.NewCategoryUsecase(mockCategoryRepository)
			err := categoryUsecase.Insert(context.Background(), test.data)
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
			data: fixture.Category,
		},
		{
			name:      "failed",
			wantError: true,
			err:       errors.New("fail when update data"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockCategoryRepository := mocks.NewCategoryRepository(t)
			mockCategoryRepository.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("entity.Category")).Return(test.err)

			categoryUsecase := usecase.NewCategoryUsecase(mockCategoryRepository)
			err := categoryUsecase.Update(context.Background(), test.slug, test.data)
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
			wantError: true,
			err:       errors.New("fail when delete data"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockCategoryRepository := mocks.NewCategoryRepository(t)
			mockCategoryRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(test.err)

			categoryUsecase := usecase.NewCategoryUsecase(mockCategoryRepository)
			err := categoryUsecase.Delete(context.Background(), test.slug)
			if err != nil && test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
