package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/soerjadi/exam/category"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/utils"
)

type pgCategoryRepository struct {
	Conn *sql.DB
}

var logger = utils.LogBuilder(true)

// NewPGCategoryRepository is bridge to create an object from product.Repository interface
func NewPGCategoryRepository(Conn *sql.DB) category.Repository {
	return &pgCategoryRepository{Conn}
}

func (p *pgCategoryRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Category, error) {
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

	result := make([]*models.Category, 0)
	for rows.Next() {
		t := new(models.Category)

		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.ParentID,
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

func (p *pgCategoryRepository) fetchRaw(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
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

func (p *pgCategoryRepository) fetchRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
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

func (p *pgCategoryRepository) GetByID(ctx context.Context, id int64) (product *models.Category, err error) {
	query := `SELECT id, name, parent_id, created, updated FROM categories WHERE id = ?`

	categories, err := p.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(categories) > 0 {
		product = categories[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (p *pgCategoryRepository) Create(ctx context.Context, category *models.Category) error {
	query := `INSERT INTO categories(name, parent_id) VALUES(?, ?) returning id`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, category.Name, category.ParentID)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = lastID
	return nil
}

func (p *pgCategoryRepository) Update(ctx context.Context, category *models.Category) error {
	query := "UPDATE categories SET name = ?, parent_id = ?, updated = ? WHERE id = ?"

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, category.Name, category.ParentID, time.Now(), category.ID)
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

func (p *pgCategoryRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM categories WHERE id = ?"

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
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

func (p *pgCategoryRepository) Search(ctx context.Context, query *string, offset int64, limit int64) ([]*models.Category, int64, error) {
	var searchQuery string

	if query != nil && len(*query) > 0 {
		searchQuery = "WHERE LOWER(name) LIKE '%?%'"
	}

	q := fmt.Sprintf("SELECT id, name, parent_id, created, updated FROM categories %s ORDER BY created LIMIT ? OFFSET ?", searchQuery)
	qCount := fmt.Sprintf("SELECT count(id) FROM categories %s", searchQuery)

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
