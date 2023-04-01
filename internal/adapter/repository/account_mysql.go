package repository

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/omid-h70/bank-service/internal/adapter/logger"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/helper"
	"github.com/pkg/errors"
)

var (
	ErrAccountRepoDBInvalidTransactionExp = errors.New("invalid NULL Transaction Exception")
	ErrAccountRepoDBWhileUpdatingBalance  = errors.New("Updating Balance Query Failed")
)

type AccountRepositoryMySqlDB struct {
	client *sql.DB
}

func (s *AccountRepositoryMySqlDB) GetAccountInfoByCard(ctx context.Context, cardNum string, mode int) (domain.AccountInfoOutput, error) {

	var (
		query string = `SELECT card.card_id, card.account_id, account.account_balance, account_rule.min_amount, account_rule.max_amount, account_rule.transaction_fee, customer.phone_number
			FROM card
			LEFT JOIN account
			ON card.account_id = account.account_id
			LEFT JOIN account_rule
			ON account.account_rule_id=account_rule.account_rule_id
			LEFT JOIN customer
			ON customer.customer_id = account.customer_id
			WHERE card.card_number=?`

		//result  sql.Result
		err     error
		balance int64
		infoOut domain.AccountInfoOutput
	)

	switch mode {
	case domain.AtomicInfo:
		{
			tx := ctx.Value("tx").(*sql.Tx)
			if tx == nil {
				return infoOut, ErrAccountRepoDBInvalidTransactionExp
			}

			err = tx.QueryRowContext(ctx, query, cardNum).Scan(
				&infoOut.CardId,
				&infoOut.AccountID,
				&infoOut.Balance,
				&infoOut.AccountRuleInfo.MinAmount,
				&infoOut.AccountRuleInfo.MaxAmount,
				&infoOut.AccountRuleInfo.TransactionFee,
				&infoOut.CustomerInfo.PhoneNum,
			)
		}
	case domain.NormalInfo:
		_, err = s.client.Exec(query, cardNum)
		if err != nil {
			err = s.client.QueryRowContext(ctx, query, cardNum).Scan(&balance)
		}
	}

	return infoOut, err
}

func (s *AccountRepositoryMySqlDB) GetBalanceByAccount(ctx context.Context, accountId string, mode int) (int64, error) {

	var (
		query   string = "SELECT account_balance FROM account WHERE  account_id=?"
		result  sql.Result
		err     error
		balance int64
	)

	switch mode {
	case domain.AtomicInfo:
		{
			tx := ctx.Value("tx").(*sql.Tx)
			if tx == nil {
				return 0, ErrAccountRepoDBInvalidTransactionExp
			}
			err = tx.QueryRowContext(ctx, query, accountId).Scan(&balance)
		}
	case domain.NormalInfo:
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

func (s *AccountRepositoryMySqlDB) UpdateBalance(ctx context.Context, accountId string, val int64) error {
	var (
		query string = "UPDATE account SET account_balance=? WHERE account_id=?"
		err   error
	)

	tx := ctx.Value("tx").(*sql.Tx)
	if tx == nil {
		return ErrAccountRepoDBInvalidTransactionExp
	}

	_, err = tx.ExecContext(ctx, query, val, accountId)
	if err != nil {
		logger.LOG(err)
		return ErrAccountRepoDBWhileUpdatingBalance
	}
	return nil
}

func NewAccountRepositoryMySqlDB(clientIn *sql.DB) AccountRepositoryMySqlDB {
	return AccountRepositoryMySqlDB{clientIn}
}
