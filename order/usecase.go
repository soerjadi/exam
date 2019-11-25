package order

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Usecase represent the order usecase
type Usecase interface {
	GetList(ctx context.Context, offset int64, limit int64) ([]*models.Order, int64, error)
	Create(ctx context.Context, order *models.Order) error
	Delete(ctx context.Context, id int64) error
}
