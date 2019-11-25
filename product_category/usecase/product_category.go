package usecase

import (
	"context"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product"
	cat "github.com/soerjadi/exam/product_category"
)

type pcUsecase struct {
	productRepo    product.Repository
	pcRepo         cat.Repository
	contextTimeout time.Duration
}

// NewPCUsecase will create new product category usecase object representation of category usecase interface
func NewPCUsecase(pc cat.Repository, timeout time.Duration) cat.Usecase {
	return &pcUsecase{
		pcRepo:         pc,
		contextTimeout: timeout,
	}
}

func (pc *pcUsecase) GetByProductID(ctx context.Context, productID int64) ([]*models.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, pc.contextTimeout)
	defer cancel()

	cats, err := pc.pcRepo.GetByProductID(ctx, productID)

	if err != nil {
		return nil, err
	}

	return cats, err
}

func (pc *pcUsecase) GetByCategoryID(ctx context.Context, categoryID int64) ([]*models.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, pc.contextTimeout)
	defer cancel()

	products, err := pc.pcRepo.GetByCategoryID(ctx, categoryID)

	if err != nil {
		return nil, err
	}

	return products, err
}

func (pc *pcUsecase) Create(ctx context.Context, cat *models.ProductCategory) error {
	ctx, cancel := context.WithTimeout(ctx, pc.contextTimeout)
	defer cancel()

	err := pc.pcRepo.Create(ctx, cat)
	if err != nil {
		return err
	}

	return nil
}

func (pc *pcUsecase) DeleteByProductID(ctx context.Context, productID int64) error {
	ctx, cancel := context.WithTimeout(ctx, pc.contextTimeout)
	defer cancel()

	return pc.pcRepo.DeleteByProductID(ctx, productID)
}

func (pc *pcUsecase) DeleteByCategoryID(ctx context.Context, categoryID int64) error {
	ctx, cancel := context.WithTimeout(ctx, pc.contextTimeout)
	defer cancel()

	return pc.pcRepo.DeleteByCategoryID(ctx, categoryID)
}
