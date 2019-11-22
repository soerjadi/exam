package repository_test

import (
	"context"
	"strings"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product/repository"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockProduct := []models.Product{
		models.Product{
			ID:      1,
			Name:    "product 1",
			SKU:     "sku 1",
			Created: time.Now(),
		},
		models.Product{
			ID:      2,
			Name:    "product 2",
			SKU:     "sku 2",
			Created: time.Now(),
		},
	}

	found := int64(2)
	rows := sqlmock.NewRows([]string{"id", "name", "sku", "created", "updated"}).
		AddRow(mockProduct[0].ID, mockProduct[0].Name, mockProduct[0].SKU, mockProduct[0].Created, mockProduct[0].Updated).
		AddRow(mockProduct[1].ID, mockProduct[1].Name, mockProduct[1].SKU, mockProduct[1].Created, mockProduct[1].Updated)

	rowCount := sqlmock.NewRows([]string{"count"}).
		AddRow(found)

	query := "SELECT id, name, sku, created, updated FROM products WHERE LOWER\\(name\\) LIKE '\\%\\?\\%' OR LOWER\\(sku\\) LIKE '\\%\\?\\%' ORDER BY created LIMIT \\? OFFSET \\?"
	countQuery := "SELECT count\\(id\\) FROM products WHERE LOWER\\(name\\) LIKE '\\%\\?\\%' OR LOWER\\(sku\\) LIKE '\\%\\?\\%'"
	searchQuery := strings.ToLower("product")

	mock.ExpectQuery(query).WithArgs(searchQuery, searchQuery, 10, 0).WillReturnRows(rows)
	mock.ExpectPrepare(countQuery).ExpectQuery().WithArgs(searchQuery, searchQuery).WillReturnRows(rowCount)

	p := repository.NewPGProductRepository(db)
	result, count, err := p.Search(context.TODO(), &searchQuery, int64(0), int64(10))

	assert.NoError(t, err)
	assert.Equal(t, found, count)
	assert.Len(t, result, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "sku", "created", "updated"}).
		AddRow(1, "product 1", "sku", time.Now(), time.Now())

	query := "SELECT id, name, sku, created, updated FROM products WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := repository.NewPGProductRepository(db)

	product, err := p.GetByID(context.TODO(), int64(1))

	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	product := &models.Product{
		Name: "product",
		SKU:  "sku",
	}

	query := "INSERT INTO products\\(name, sku\\) VALUES\\(\\?, \\?\\) returning id"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(product.Name, product.SKU).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGProductRepository(db)

	err = p.Create(context.TODO(), product)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), product.ID)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	product := &models.Product{
		ID:      2,
		Name:    "product",
		SKU:     "sku",
		Created: time.Now(),
	}

	query := "UPDATE product SET name = \\?, sku = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(product.Name, product.SKU, product.ID).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGProductRepository(db)

	err = p.Update(context.TODO(), product)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM products WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(2).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGProductRepository(db)

	err = p.Delete(context.TODO(), int64(2))
	assert.NoError(t, err)
}
