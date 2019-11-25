package product_category

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Usecase represent the product category usecase
type Usecase interface {
	GetByProductID(ctx context.Context, productID int64) ([]*models.ProductCategory, error)
	GetByCategoryID(ctx context.Context, categoryID int64) ([]*models.ProductCategory, error)
	Create(ctx context.Context, pc *models.ProductCategory) error
	DeleteByProductID(ctx context.Context, productID int64) error
	DeleteByCategoryID(ctx context.Context, categoryID int64) error
}
