package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/soerjadi/exam/utils"
)

// RDB initialize
func RDB() Database {
	return &Impl{}
}

// DB implementation database connection
func (d Impl) DB() *sql.DB {
	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/soerjadi/gopos/.env")); err != nil {
		panic("Error loading .env file")
	}

	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbUser := utils.GetEnv("DB_USER", "")
	dbPort := utils.GetEnv("DB_PORT", "")
	dbPass := utils.GetEnv("DB_PASS", "")
	dbName := utils.GetEnv("DB_NAME", "")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(1 * time.Second)

	return db
}

// DBTest implementation database test connection
func (d Impl) DBTest() *sql.DB {
	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/soerjadi/gopos/.env")); err != nil {
		panic("Error loading .env file")
	}

	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbUser := utils.GetEnv("DB_USER", "")
	dbPort := utils.GetEnv("DB_PORT", "")
	dbPass := utils.GetEnv("DB_PASS", "")
	dbName := utils.GetEnv("DB_TEST", "")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	dbTest, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = dbTest.Ping()
	if err != nil {
		panic(err)
	}

	return dbTest
}

// DBTestRepository repository for database Test
func DBTestRepository(conn *sql.DB) DBTest {
	return &DBTestImpl{Conn: conn}
}

func (db *DBTestImpl) truncate(table string) error {
	cmd := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)

	_, err := db.Conn.Exec(cmd)
	return err
}

// Clean database test table
func (db *DBTestImpl) Clean(tables ...string) {
	for _, table := range tables {
		if err := db.truncate(table); err != nil {
			panic(err)
		}
	}
}
