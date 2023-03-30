package repository

import (
	"database/sql"
	"fmt"
	"time"
)

func NewRepositoryMySqlDB(dbName string) *sql.DB {
	client, err := sql.Open("mysql", fmt.Sprintf("admin:admin@tcp/%s", dbName))
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
