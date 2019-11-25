package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product"
	"github.com/soerjadi/exam/utils"
)

type pgProductRepository struct {
	Conn *sql.DB
}

var logger = utils.LogBuilder(true)

// NewPGProductRepository is bridge to create an object from product.Repository interface
func NewPGProductRepository(Conn *sql.DB) product.Repository {
	return &pgProductRepository{Conn}
}

func (p *pgProductRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Product, error) {
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

	result := make([]*models.Product, 0)
	for rows.Next() {
		t := new(models.Product)

		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.SKU,
			&t.Created,
			&t.Updated,
		)

		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (p *pgProductRepository) fetchRaw(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
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

	return rows, nil
}

func (p *pgProductRepository) fetchRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	stmt, err := p.Conn.PrepareContext(ctx, query)

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

func (p *pgProductRepository) GetByID(ctx context.Context, id int64) (product *models.Product, err error) {
	query := `SELECT id, name, sku, created, updated FROM products WHERE id = ?`

	products, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(products) > 0 {
		product = products[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (p *pgProductRepository) Create(ctx context.Context, product *models.Product) error {
	query := `INSERT INTO products(name, sku) VALUES(?, ?) returning id`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, product.Name, product.SKU)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = lastID
	return nil
}

func (p *pgProductRepository) Update(ctx context.Context, product *models.Product) error {
	query := "UPDATE products SET name = ?, sku = ?, updated = ? WHERE id = ?"

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.SKU, time.Now(), product.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("Total affected: %d", affected)
	}

	return nil

}

func (p *pgProductRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM products WHERE id = ?"

	stmt, err := p.Conn.PrepareContext(ctx, query)
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
		err = fmt.Errorf("Total affected: %d", rowsAffected)
		return err
	}

	return nil
}

func (p *pgProductRepository) Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Product, int64, error) {
	var searchQuery string

	if query != nil && len(*query) > 0 {
		searchQuery = "WHERE LOWER(name) LIKE '%?%' OR LOWER(sku) LIKE '%?%'"
	}

	q := fmt.Sprintf("SELECT id, name, sku, created, updated FROM products %s ORDER BY created LIMIT ? OFFSET ?", searchQuery)
	qCount := fmt.Sprintf("SELECT count(id) FROM products %s", searchQuery)

	result, err := p.fetch(ctx, q, strings.ToLower(*query), strings.ToLower(*query), limit, offset)

	if err != nil {
		return nil, 0, err
	}

	rows, err := p.fetchRow(ctx, qCount, strings.ToLower(*query), strings.ToLower(*query))

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
