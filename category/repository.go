package category

import (
	"context"

	"github.com/soerjadi/exam/models"
)

// Repository represent the category repository
type Repository interface {
	Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Category, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Category, error)
	Create(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int64) error
}
