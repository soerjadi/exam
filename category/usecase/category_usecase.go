package usecase

import (
	"context"
	"time"

	"github.com/soerjadi/exam/category"
	"github.com/soerjadi/exam/models"
	"gopkg.in/guregu/null.v3"
)

type categoryUsecase struct {
	repo           category.Repository
	contextTimeout time.Duration
}

// NewCategoryUsecase will create object that represent of category.Usecase interface
func NewCategoryUsecase(c category.Repository, timeout time.Duration) category.Usecase {
	return &categoryUsecase{
		repo:           c,
		contextTimeout: timeout,
	}
}

func (c *categoryUsecase) Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Category, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categories, found, err := c.repo.Search(ctx, query, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return categories, found, nil
}

func (c *categoryUsecase) GetByID(ctx context.Context, id int64) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	category, err := c.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryUsecase) Create(ctx context.Context, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.repo.Create(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryUsecase) Update(ctx context.Context, category *models.Category) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	category.Updated = null.NewTime(
		time.Now(), true,
	)

	return c.repo.Update(ctx, category)
}

func (c *categoryUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	exists, err := c.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if exists == nil {
		return models.ErrNotFound
	}

	return c.repo.Delete(ctx, id)
}
