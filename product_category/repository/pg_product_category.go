package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/soerjadi/exam/models"
	cat "github.com/soerjadi/exam/product_category"
	"github.com/soerjadi/exam/utils"
)

type pgProductCategoryRepository struct {
	Conn *sql.DB
}

var logger = utils.LogBuilder(true)

// NewPGProductCategoryRepository is bridge to create an object from productCategory.Repository interface
func NewPGProductCategoryRepository(Conn *sql.DB) cat.Repository {
	return &pgProductCategoryRepository{Conn}
}

func (p *pgProductCategoryRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.ProductCategory, error) {
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

	result := make([]*models.ProductCategory, 0)
	for rows.Next() {
		t := new(models.ProductCategory)

		err = rows.Scan(
			&t.ID,
			&t.ProductID,
			&t.CategoryID,
		)

		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (p *pgProductCategoryRepository) GetByProductID(ctx context.Context, id int64) ([]*models.ProductCategory, error) {
	query := `SELECT id, product_id, category_id FROM product_category WHERE product_id = ?`

	cats, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return cats, nil
}

func (p *pgProductCategoryRepository) GetByCategoryID(ctx context.Context, id int64) ([]*models.ProductCategory, error) {
	query := `SELECT id, product_id, category_id FROM product_category WHERE category_id = ?`

	products, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *pgProductCategoryRepository) Create(ctx context.Context, pc *models.ProductCategory) error {
	query := `INSERT INTO product_category(product_id, category_id) VALUES(?, ?) returning id`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, pc.ProductID, pc.CategoryID)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	pc.ID = lastID
	return nil
}

func (p *pgProductCategoryRepository) DeleteByProductID(ctx context.Context, productID int64) error {
	query := `DELETE FROM product_category WHERE product_id = ?`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, productID)
	if err != nil {
		return nil
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

func (p *pgProductCategoryRepository) DeleteByCategoryID(ctx context.Context, categoryID int64) error {
	query := `DELETE FROM product_category WHERE category_id = ?`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, categoryID)
	if err != nil {
		return nil
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
