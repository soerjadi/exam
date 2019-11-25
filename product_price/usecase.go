package product_price

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Usecase represent the product price usecase
type Usecase interface {
	Create(ctx context.Context, price *models.ProductPrice) error
	DeleteByProductID(ctx context.Context, id int64) error
	GetByProductID(ctx context.Context, id int64) ([]*models.ProductPrice, error)
	GetPriceByAmount(ctx context.Context, amount int64) (*models.ProductPrice, error)
}
