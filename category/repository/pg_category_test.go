package repository_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/soerjadi/exam/category/repository"
	"github.com/soerjadi/exam/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
)

func TestSearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockCategory := []models.Category{
		models.Category{
			ID:       1,
			Name:     "category 1",
			ParentID: null.NewInt(int64(0), true),
			Created:  time.Now(),
		},
		models.Category{
			ID:       2,
			Name:     "category 2",
			ParentID: null.NewInt(int64(1), true),
			Created:  time.Now(),
		},
	}

	found := int64(2)
	rows := sqlmock.NewRows([]string{"id", "name", "parent_id", "created", "updated"}).
		AddRow(mockCategory[0].ID, mockCategory[0].Name, mockCategory[0].ParentID, mockCategory[0].Created, mockCategory[0].Updated).
		AddRow(mockCategory[1].ID, mockCategory[1].Name, mockCategory[1].ParentID, mockCategory[1].Created, mockCategory[1].Updated)

	rowCount := sqlmock.NewRows([]string{"count"}).
		AddRow(found)

	query := "SELECT id, name, parent_id, created, updated FROM categories WHERE LOWER\\(name\\) LIKE '\\%\\?\\%' ORDER BY created LIMIT \\? OFFSET \\?"
	countQuery := "SELECT count\\(id\\) FROM categories WHERE LOWER\\(name\\) LIKE '\\%\\?\\%'"
	searchQuery := strings.ToLower("category")

	mock.ExpectQuery(query).WithArgs(searchQuery, searchQuery, 10, 0).WillReturnRows(rows)
	mock.ExpectPrepare(countQuery).ExpectQuery().WithArgs(searchQuery, searchQuery).WillReturnRows(rowCount)

	p := repository.NewPGCategoryRepository(db)
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

	rows := sqlmock.NewRows([]string{"id", "name", "parent_id", "created", "updated"}).
		AddRow(1, "category 1", 0, time.Now(), time.Now())

	query := "SELECT id, name, parent_id, created, updated FROM categories WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := repository.NewPGCategoryRepository(db)

	category, err := p.GetByID(context.TODO(), int64(1))

	assert.NoError(t, err)
	assert.NotNil(t, category)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	category := &models.Category{
		Name:     "category",
		ParentID: null.NewInt(int64(0), true),
	}

	query := "INSERT INTO categories\\(name, parent_id\\) VALUES\\(\\?, \\?\\) returning id"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(category.Name, category.ParentID).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGCategoryRepository(db)

	err = p.Create(context.TODO(), category)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), category.ID)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	category := &models.Category{
		ID:       2,
		Name:     "category",
		ParentID: null.NewInt(int64(0), true),
		Created:  time.Now(),
	}

	query := "UPDATE categories SET name = \\?, parent_id = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(category.Name, category.ParentID, category.ID).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGCategoryRepository(db)

	err = p.Update(context.TODO(), category)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM categories WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(2).WillReturnResult(sqlmock.NewResult(2, 1))

	p := repository.NewPGCategoryRepository(db)

	err = p.Delete(context.TODO(), int64(2))
	assert.NoError(t, err)
}
