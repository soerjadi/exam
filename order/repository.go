package order

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Repository represent the order repository interface
type Repository interface {
	GetList(ctx context.Context, offset int64, limit int64) ([]*models.Order, int64, error)
	Create(ctx context.Context, order *models.Order) error
	Delete(ctx context.Context, id int64) error
}
