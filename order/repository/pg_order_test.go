package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/soerjadi/exam/models"
	"github.com/soerjadi/exam/order/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockOrder := []*models.Order{
		&models.Order{
			ID:        int64(8),
			ProductID: int64(2),
			Amount:    int64(20), // with amount 20 -> 8000
			Price:     160000.0,
			Status:    models.OrderProccessed,
			Created:   time.Now(),
		},
		&models.Order{
			ID:        int64(9),
			ProductID: int64(3),
			Amount:    int64(1),
			Price:     10000.0,
			Status:    models.OrderShipped,
			Created:   time.Now(),
		},
	}

	found := int64(2)
	rows := sqlmock.NewRows([]string{"id", "product_id", "amount", "price", "status", "created"}).
		AddRow(mockOrder[0].ID, mockOrder[0].ProductID, mockOrder[0].Amount, mockOrder[0].Price, mockOrder[0].Status, mockOrder[0].Created).
		AddRow(mockOrder[0].ID, mockOrder[1].ProductID, mockOrder[1].Amount, mockOrder[1].Price, mockOrder[1].Status, mockOrder[1].Created)

	rowCount := sqlmock.NewRows([]string{"count"}).AddRow(found)

	query := "SELECT id, product_id, amount, price, status, created FROM orders ORDER BY created OFFSET \\? LIMIT \\?"
	cQuery := "SELECT count\\(id\\) FROM orders"

	mock.ExpectQuery(query).WithArgs(int64(0), int64(10)).WillReturnRows(rows)
	mock.ExpectPrepare(cQuery).ExpectQuery().WillReturnRows(rowCount)

	p := repository.NewPGOrderRepository(db)
	result, count, err := p.GetList(context.TODO(), int64(0), int64(10))

	assert.NoError(t, err)
	assert.Equal(t, found, count)
	assert.Len(t, result, 2)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	order := &models.Order{
		ProductID: int64(2),
		Amount:    int64(20), // with amount 20 -> 8000
		Price:     160000.0,
		Status:    models.OrderProccessed,
	}

	query := "INSERT INTO orders\\(product_id, amount, price, status\\) VALUES\\(\\?, \\?, \\?, \\?\\) returning id"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(order.ProductID, order.Amount, order.Price, order.Status).WillReturnResult(sqlmock.NewResult(89, 1))

	p := repository.NewPGOrderRepository(db)

	err = p.Create(context.TODO(), order)

	assert.NoError(t, err)
	assert.Equal(t, int64(89), order.ID)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM orders WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(int64(9)).WillReturnResult(sqlmock.NewResult(9, 1))

	p := repository.NewPGOrderRepository(db)

	err = p.Delete(context.TODO(), int64(9))
	assert.NoError(t, err)
}
