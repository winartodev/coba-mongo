package usecase

import (
	"context"
	"winartodev/coba-mongodb/entity"
	"winartodev/coba-mongodb/repository"
)

type ProductUsecase interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	Insert(ctx context.Context, product entity.Product) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (p *productUsecase) FindAll(ctx context.Context) ([]entity.Product, error) {
	res, err := p.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *productUsecase) Insert(ctx context.Context, product entity.Product) error {
	err := p.repo.Insert(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
