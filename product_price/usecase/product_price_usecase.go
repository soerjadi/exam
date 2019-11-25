package usecase

import (
	"context"
	"time"

	"github.com/soerjadi/exam/models"
	productPrice "github.com/soerjadi/exam/product_price"
)

type productPriceUsecase struct {
	repo           productPrice.Repository
	contextTimeout time.Duration
}

// NewProductPriceUsecase will create object that represent of product price usecase interface
func NewProductPriceUsecase(p productPrice.Repository, timeout time.Duration) productPrice.Usecase {
	return &productPriceUsecase{
		repo:           p,
		contextTimeout: timeout,
	}
}

func (p *productPriceUsecase) Create(ctx context.Context, price *models.ProductPrice) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err := p.repo.Create(ctx, price)
	if err != nil {
		return err
	}

	return nil
}

func (p *productPriceUsecase) DeleteByProductID(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err := p.repo.DeleteByProductID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *productPriceUsecase) GetByProductID(ctx context.Context, id int64) ([]*models.ProductPrice, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	prices, err := p.repo.GetByProductID(ctx, id)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (p *productPriceUsecase) GetPriceByAmount(ctx context.Context, amount int64) (*models.ProductPrice, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	price, err := p.repo.GetPriceByAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	return price, nil
}
