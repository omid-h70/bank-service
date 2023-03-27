package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

//func (s CustomerRepositoryDB) connectToDb() (*sql.DB, error) {
//	client, err := sql.Open("data", "user:password@/dbname")
//	if err != nil {
//		panic(err)
//	}
//	client.SetConnMaxLifetime(time.Minute * 3)
//	client.SetMaxOpenConns(10)
//	client.SetMaxIdleConns(10)
//	return client, nil
//}

func (s CustomerRepositoryDB) FindAll() (customer []domain.Customer, err error) {
	findAllSql := "select *"
	//client, err := s.connectToDb()
	//if err != nil {
	//	return nil, err
	//}

	rows, err := s.client.Query(findAllSql)
	if err != nil {
		return nil, err
	}

	customers := make([]domain.Customer, 0)
	for rows.Next() {
		var c domain.Customer
		err := rows.Scan(c.Id)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryMySqlDB() CustomerRepositoryDB {
	client, err := sql.Open("data", "user:pass@tcp/dbname")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
