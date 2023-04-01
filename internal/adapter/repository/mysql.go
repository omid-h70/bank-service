package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type MySqlConfig struct {
	DbServerAddr string
	DbServerPort string
	DbName       string
	DbUser       string
	DbPass       string
}

func NewRepositoryMySqlDB(c MySqlConfig) *sql.DB {

	var (
		client *sql.DB
		err    error
	)

	for i := 0; i < 10; i++ {
		client, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			c.DbUser,
			c.DbPass,
			c.DbServerAddr,
			c.DbServerPort,
			c.DbName))

		if err != nil {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Sleep A little...")

		} else {
			if err = client.Ping(); err != nil {
				continue
			}
			fmt.Println("Connected in " + fmt.Sprintf("%d", i) + " Attempt")
			//client.Query("USE webServiceDB")
			break
		}
	}

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
