package usecase

import (
	"context"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product"
	"gopkg.in/guregu/null.v3"
)

type productUsecase struct {
	repo           product.Repository
	contextTimeout time.Duration
}

// NewProductUsecase will create object that represent of product.Usecase interface
func NewProductUsecase(p product.Repository, timeout time.Duration) product.Usecase {
	return &productUsecase{
		repo:           p,
		contextTimeout: timeout,
	}
}

func (p *productUsecase) Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Product, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	products, found, err := p.repo.Search(ctx, query, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return products, found, nil
}

func (p *productUsecase) GetByID(ctx context.Context, id int64) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	product, err := p.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productUsecase) Create(ctx context.Context, product *models.Product) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err := p.repo.Create(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p *productUsecase) Update(ctx context.Context, product *models.Product) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	product.Updated = null.NewTime(
		time.Now(), true,
	)

	return p.repo.Update(ctx, product)
}

func (p *productUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	exists, err := p.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if exists == nil {
		return models.ErrNotFound
	}

	return p.repo.Delete(ctx, id)
}

func (p *productUsecase) Compare(ctx context.Context, id1 int64, id2 int64) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	products := make([]*models.Product, 0)
	product1, err := p.repo.GetByID(ctx, id1)

	if err == nil {
		products = append(products, product1)
	}

	product2, err := p.repo.GetByID(ctx, id2)

	if err == nil {
		products = append(products, product2)
	}

	return products, nil
}
