package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/helper"
	"github.com/pkg/errors"
	"time"
)

type AccountRepositoryDBIMpl struct {
	client *sql.DB
	tx     *sql.Tx
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

func (s AccountRepositoryDBIMpl) FindAll() (customer []domain.Customer, err error) {
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

func (s AccountRepositoryDBIMpl) MakeTransferFromCardToCard(ctx context.Context, input domain.CardTransferInput) (domain.CardTransferOutput, error) {

	var (
		//query  string = "UPDATE account SET account_balance=? WHERE account_id=?"
		//result sql.Result
		err error
	)

	tx, err := s.client.BeginTx(ctx, nil)
	defer func() {
		err := tx.Rollback()
		fmt.Println(err)
	}()

	if err != nil {
		return domain.CardTransferOutput{}, errors.Wrap(err, "error updating account balance")
	}

	//---------------Get Atomic balance
	s.tx = tx
	balanceFrom, errFrom := s.GetBalanceByCard(ctx, input.CardFrom.CardNum, domain.AtomicBalance)
	balanceTo, errTo := s.GetBalanceByCard(ctx, input.CardTo.CardNum, domain.AtomicBalance)
	if errFrom != nil && errTo != nil {
		return domain.CardTransferOutput{}, errors.Wrap(err, "error updating account balance")
	}

	//---------------Calculations
	balanceFrom -= input.Amount
	balanceTo += input.Amount

	resultFrom := s.UpdateBalance(ctx, input.CardFrom.AccountId, balanceTo)
	resultTo := s.UpdateBalance(ctx, input.CardTo.AccountId, balanceFrom)
	if resultFrom != nil || resultTo != nil {
		return domain.CardTransferOutput{}, errors.Wrap(err, "error updating account balance")
	}

	//---------------Commit Transaction
	if err = tx.Commit(); err != nil {
		//return fail(err)
	}

	return domain.CardTransferOutput{}, nil
}

func (s AccountRepositoryDBIMpl) InsertTransaction(ctx context.Context, input domain.CardTransferInput) error {

	var (
		query  string = fmt.Sprintf("INSERT INTO transaction (VALUES account_id_from=?, account_id_to=?, transction_type=%d, amount = ?)", domain.TransactionTransfer)
		result bool
		err    error
	)

	if s.tx == nil {

	}
	err = s.tx.QueryRowContext(ctx, query, input.CardFrom.CardNum, input.CardTo.CardNum, input.Amount).Scan(&result)
	//if err != nil {
	//
	//}

	helper.GO_UNUSED(result, err)
	return nil
}

func (s AccountRepositoryDBIMpl) GetBalanceByCard(ctx context.Context, cardNum string, mode int) (int64, error) {

	var (
		query string = `SELECT card.account_id, account.account_balance 
						  FROM card LEFT JOIN account 
   						  ON card.account_id = account.account_id 
						  WHERE card.card_number=?`

		result  sql.Result
		err     error
		balance int64
	)

	switch mode {
	case domain.AtomicBalance:
		{
			if s.tx == nil {

			}
			err = s.tx.QueryRowContext(ctx, query, cardNum).Scan(&balance)
		}
	case domain.NormalBalance:
		result, err = s.client.Exec(query, cardNum)
		if err != nil {
			err = s.client.QueryRowContext(ctx, query, cardNum).Scan(&balance)
		}
	}

	//if err != nil {
	//
	//}
	helper.GO_UNUSED(result, err)
	return balance, nil
}

func (s AccountRepositoryDBIMpl) GetBalanceByAccount(ctx context.Context, accountId string, mode int) (int64, error) {

	var (
		query   string = "SELECT account_balance FROM account WHERE  account_id=?"
		result  sql.Result
		err     error
		balance int64
	)

	switch mode {
	case domain.AtomicBalance:
		{
			if s.tx == nil {

			}
			err = s.tx.QueryRowContext(ctx, query, accountId).Scan(&balance)
		}
	case domain.NormalBalance:
		result, err = s.client.Exec(query, accountId)
		if err != nil {
			err = s.client.QueryRowContext(ctx, query, accountId).Scan(&balance)
		}
	}

	//if err != nil {
	//
	//}
	helper.GO_UNUSED(result, err)
	return balance, nil
}

func (s AccountRepositoryDBIMpl) UpdateBalance(ctx context.Context, accountId string, val int64) error {
	var (
		query  string = "UPDATE account SET account_balance=? WHERE account_id=?"
		result bool
		err    error
	)

	if s.tx == nil {

	}
	err = s.tx.QueryRowContext(ctx, query, val, accountId).Scan(&result)
	//if err != nil {
	//
	//}

	helper.GO_UNUSED(result, err)
	return nil
}

func NewAccountRepositoryMySqlDB(dbName string) AccountRepositoryDBIMpl {
	client, err := sql.Open("mysql", fmt.Sprintf("admin:admin@tcp/%s", dbName))
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return AccountRepositoryDBIMpl{client, nil}
}
