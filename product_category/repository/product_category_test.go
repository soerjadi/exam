package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product_category/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	pc := &models.ProductCategory{
		ProductID:  int64(9),
		CategoryID: int64(10),
	}

	query := "INSERT INTO product_category\\(product_id, category_id\\) VALUES\\(\\?, \\?\\) returning id"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(pc.ProductID, pc.CategoryID).WillReturnResult(sqlmock.NewResult(8, 1))

	p := repository.NewPGProductCategoryRepository(db)

	err = p.Create(context.TODO(), pc)

	assert.NoError(t, err)
	assert.Equal(t, int64(9), pc.ProductID)
}

func TestGetByProductID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "product_id", "category_id"}).
		AddRow(1, 8, 9).
		AddRow(2, 8, 10)

	query := "SELECT id, product_id, category_id FROM product_category WHERE product_id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := repository.NewPGProductCategoryRepository(db)

	cats, err := p.GetByProductID(context.TODO(), int64(8))

	assert.NoError(t, err)
	assert.NotNil(t, cats)
	assert.Equal(t, 2, len(cats))
}

func TestGetByCategoryID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "product_id", "category_id"}).
		AddRow(1, 8, 10).
		AddRow(7, 34, 10).
		AddRow(20, 56, 10)

	query := "SELECT id, product_id, category_id FROM product_category WHERE category_id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := repository.NewPGProductCategoryRepository(db)

	products, err := p.GetByCategoryID(context.TODO(), int64(8))

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 3, len(products))
}

func TestDeleteByCategoryID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM product_category WHERE product_id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(2).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGProductCategoryRepository(db)

	err = p.DeleteByProductID(context.TODO(), int64(2))
	assert.NoError(t, err)
}

func TestDeleteByProductID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM product_category WHERE category_id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(2).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGProductCategoryRepository(db)

	err = p.DeleteByCategoryID(context.TODO(), int64(2))
	assert.NoError(t, err)
}
