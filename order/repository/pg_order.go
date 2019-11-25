package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/order"
	"github.com/soerjadi/exam/utils"
)

type pgOrderRepository struct {
	Conn *sql.DB
}

var logger = utils.LogBuilder(true)

// NewPGOrderRepository is bridge to create an object from order.Repository interface
func NewPGOrderRepository(Conn *sql.DB) order.Repository {
	return &pgOrderRepository{Conn}
}

func (o *pgOrderRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Order, error) {
	rows, err := o.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	result := make([]*models.Order, 0)
	for rows.Next() {
		t := new(models.Order)

		err = rows.Scan(
			&t.ID,
			&t.ProductID,
			&t.Amount,
			&t.Price,
			&t.Status,
			&t.Created,
		)

		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (o *pgOrderRepository) fetchRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	stmt, err := o.Conn.PrepareContext(ctx, query)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	return stmt.QueryRow(args...), nil
}

func (o *pgOrderRepository) GetList(ctx context.Context, offset int64, limit int64) (orders []*models.Order, found int64, err error) {
	query := `SELECT id, product_id, amount, price, status, created FROM orders ORDER BY created OFFSET ? LIMIT ?`
	qCount := `SELECT count(id) FROM orders`

	result, err := o.fetch(ctx, query, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	rows, err := o.fetchRow(ctx, qCount)
	if err != nil {
		logger.Error(err)
		return nil, 0, err
	}

	var count int64
	err = rows.Scan(&count)

	if err != nil {
		logger.Error(err)
		return nil, 0, err
	}

	return result, count, nil
}

func (o *pgOrderRepository) Create(ctx context.Context, order *models.Order) error {
	query := `INSERT INTO orders(product_id, amount, price, status) VALUES(?, ?, ?, ?) returning id`

	stmt, err := o.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, order.ProductID, order.Amount, order.Price, order.Status)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	order.ID = lastID
	return nil
}

func (o *pgOrderRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM orders WHERE id = ?`

	stmt, err := o.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Internal Server Error")
		return err
	}

	return nil
}
