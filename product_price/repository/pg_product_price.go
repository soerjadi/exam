package repository

import (
	"context"
	"database/sql"

	"github.com/soerjadi/exam/models"
	price "github.com/soerjadi/exam/product_price"
	"github.com/soerjadi/exam/utils"
)

type pgProductPriceRepository struct {
	Conn *sql.DB
}

var logger = utils.LogBuilder(true)

// NewPGProductPriceRepository is bridge to create an object from price.Repository interface
func NewPGProductPriceRepository(Conn *sql.DB) price.Repository {
	return &pgProductPriceRepository{Conn}
}

func (p *pgProductPriceRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.ProductPrice, error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
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

	result := make([]*models.ProductPrice, 0)
	for rows.Next() {
		t := new(models.ProductPrice)

		err = rows.Scan(
			&t.ID,
			&t.Amount,
			&t.Price,
			&t.ProductID,
		)

		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (p *pgProductPriceRepository) Create(ctx context.Context, price *models.ProductPrice) error {
	query := `INSERT INTO product_price(amount, price, product_id) VALUES(?, ?, ?) returning id`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, price.Amount, price.Price, price.ProductID)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	price.ID = lastID
	return nil
}

func (p *pgProductPriceRepository) GetByProductID(ctx context.Context, id int64) ([]*models.ProductPrice, error) {
	query := `SELECT id, amount, price, product_id FROM product_price WHERE product_id = ?`

	prices, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (p *pgProductPriceRepository) GetPriceByAmount(ctx context.Context, amount int64) (price *models.ProductPrice, err error) {
	query := `SELECT id, amount, price, product_id FROM product_price WHERE amount < ? LIMIT 1`

	prices, err := p.fetch(ctx, query, amount)
	if err != nil {
		return nil, err
	}

	if len(prices) > 0 {
		price = prices[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (p *pgProductPriceRepository) DeleteByProductID(ctx context.Context, id int64) error {
	query := `DELETE FROM product_price WHERE product_id = ?`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
