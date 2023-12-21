package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var (
	dbType     = os.Getenv("DB_TYPE")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
)

func Setup() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	client, err := sqlx.Connect(dbType, dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
