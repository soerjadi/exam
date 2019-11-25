package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/product_price/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	price := &models.ProductPrice{
		Amount:    20,
		Price:     9000.0,
		ProductID: 8,
	}

	query := "INSERT INTO product_price\\(amount, price, product_id\\) VALUES\\(\\?, \\?, \\?\\) returning id"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(price.Amount, price.Price, price.ProductID).WillReturnResult(sqlmock.NewResult(6, 1))

	p := repository.NewPGProductPriceRepository(db)

	err = p.Create(context.TODO(), price)

	assert.NoError(t, err)
	assert.Equal(t, int64(6), price.ID)
}

func TestGetByProductID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockProductPrice := []*models.ProductPrice{
		&models.ProductPrice{
			ID:        1,
			Amount:    20,
			Price:     9000.0,
			ProductID: 8,
		},
		&models.ProductPrice{
			ID:        2,
			Amount:    30,
			Price:     8000.0,
			ProductID: 8,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "amount", "price", "product_id"}).
		AddRow(mockProductPrice[0].ID, mockProductPrice[0].Amount, mockProductPrice[0].Price, mockProductPrice[0].ProductID).
		AddRow(mockProductPrice[1].ID, mockProductPrice[1].Amount, mockProductPrice[1].Price, mockProductPrice[1].ProductID)

	query := "SELECT id, amount, price, product_id FROM product_price WHERE product_id = \\?"

	mock.ExpectQuery(query).WithArgs(int64(8)).WillReturnRows(rows)

	p := repository.NewPGProductPriceRepository(db)
	prices, err := p.GetByProductID(context.TODO(), int64(8))

	assert.NoError(t, err)
	assert.Equal(t, mockProductPrice, prices)

}

func TestGetPriceByAmount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockProductPrice := models.ProductPrice{
		Amount:    20,
		Price:     9000.0,
		ProductID: 8,
	}
	mockProductPrice2 := models.ProductPrice{
		Amount:    30,
		Price:     8000.0,
		ProductID: 8,
	}
	mockProductPrice3 := models.ProductPrice{
		Amount:    40,
		Price:     9000.0,
		ProductID: 8,
	}

	rows := sqlmock.NewRows([]string{"id", "amount", "price", "product_id"}).
		AddRow(mockProductPrice2.ID, mockProductPrice2.Amount, mockProductPrice2.Price, mockProductPrice2.ProductID)

	query := "SELECT id, amount, price, product_id FROM product_price WHERE amount < \\?"

	mock.ExpectQuery(query).WithArgs(int64(35)).WillReturnRows(rows)

	p := repository.NewPGProductPriceRepository(db)

	product, err := p.GetPriceByAmount(context.TODO(), int64(35))

	assert.NoError(t, err)
	assert.Equal(t, &mockProductPrice2, product)
	assert.NotEqual(t, &mockProductPrice, product)
	assert.NotEqual(t, &mockProductPrice3, product)

}

func TestDeleteByProductID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM product_price WHERE product_id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(2).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGProductPriceRepository(db)

	err = p.DeleteByProductID(context.TODO(), int64(2))
	assert.NoError(t, err)
}
