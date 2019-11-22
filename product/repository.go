package product

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Repository represent the product's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (product *models.Product, err error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int64) error
	Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Product, int64, error)
}
