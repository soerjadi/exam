package usecase

import (
	"context"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/order"
)

type orderUsecase struct {
	repo    order.Repository
	timeout time.Duration
}

// NewOrderUsecase will create object that represent of order.Usecase interface
func NewOrderUsecase(o order.Repository, timeout time.Duration) order.Usecase {
	return &orderUsecase{
		repo:    o,
		timeout: timeout,
	}
}

func (o *orderUsecase) GetList(ctx context.Context, offset int64, limit int64) ([]*models.Order, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, o.timeout)
	defer cancel()

	orders, found, err := o.repo.GetList(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return orders, found, nil
}

func (o *orderUsecase) Create(ctx context.Context, order *models.Order) error {
	ctx, cancel := context.WithTimeout(ctx, o.timeout)
	defer cancel()

	err := o.repo.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, o.timeout)
	defer cancel()

	return o.repo.Delete(ctx, id)
}
