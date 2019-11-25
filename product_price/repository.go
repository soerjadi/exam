package product_price

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Repository represent the product price repository contract
type Repository interface {
	Create(ctx context.Context, price *models.ProductPrice) error
	DeleteByProductID(ctx context.Context, id int64) error
	GetByProductID(ctx context.Context, id int64) ([]*models.ProductPrice, error)
	GetPriceByAmount(ctx context.Context, amount int64) (price *models.ProductPrice, err error)
}
