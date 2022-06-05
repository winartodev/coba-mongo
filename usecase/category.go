package usecase

import (
	"context"
	"winartodev/coba-mongodb/entity"
	"winartodev/coba-mongodb/repository"
)

type CategoryUsecase interface {
	FindAll(ctx context.Context) ([]entity.Category, error)
	FindOne(ctx context.Context, slug string) (*entity.Category, error)
	Insert(ctx context.Context, category entity.Category) error
	Update(ctx context.Context, slug string, category entity.Category) error
	Delete(ctx context.Context, slug string) error
}

type category struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &category{repo: repo}
}

func (c *category) FindAll(ctx context.Context) ([]entity.Category, error) {
	res, err := c.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *category) FindOne(ctx context.Context, slug string) (*entity.Category, error) {
	res, err := c.repo.FindOne(ctx, slug)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *category) Insert(ctx context.Context, category entity.Category) error {
	err := c.repo.Insert(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (c *category) Update(ctx context.Context, slug string, category entity.Category) error {
	err := c.repo.Update(ctx, slug, category)
	if err != nil {
		return err
	}

	return nil
}

func (c *category) Delete(ctx context.Context, slug string) error {
	err := c.repo.Delete(ctx, slug)
	if err != nil {
		return err
	}

	return nil
}
