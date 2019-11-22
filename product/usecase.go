package product

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Usecase represent the product's usecase
type Usecase interface {
	Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Product, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int64) error
	Compare(ctx context.Context, id1 int64, id2 int64) ([]*models.Product, error)
}
