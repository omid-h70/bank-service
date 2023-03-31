package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/pkg/errors"
)

var (
	ErrTransRepoDBInvalidTransactionExp     = errors.New("invalid NULL Transaction Exception")
	ErrTransRepoDBUpdatingBalanceFailed     = errors.New("Error while updating Balance")
	ErrTransRepoDBWhileFetchingAccountInfo  = errors.New("Error while Fetching Account Info")
	ErrTransRepoDBWhileInsertingTransaction = errors.New("Error In Transaction")
)

type TransactionRepositoryMySqlDB struct {
	client *sql.DB
}

func (t *TransactionRepositoryMySqlDB) MakeTransferFromCardToCard(ctx context.Context, input domain.Transaction) ([]domain.AccountInfoOutput, error) {

	var (
		cardFrom        = input.CardFrom()
		cardTo          = input.CardTo()
		err             error
		accountInfoList = make([]domain.AccountInfoOutput, 2)
	)

	tx, err := t.client.BeginTx(ctx, nil)
	if err != nil {
		return accountInfoList, errors.Wrap(err, "error updating account balance")
	}
	defer func() {
		err := tx.Rollback()
		fmt.Println(err)
	}()

	//---------------Get Atomic balance
	ctx = context.WithValue(ctx, "tx", tx)

	accountRepoHandle := NewAccountRepositoryMySqlDB(t.client)
	cardFromInfo, errFrom := accountRepoHandle.GetAccountInfoByCard(ctx, cardFrom.CardNum(), domain.AtomicInfo)
	cardToInfo, errTo := accountRepoHandle.GetAccountInfoByCard(ctx, cardTo.CardNum(), domain.AtomicInfo)
	if errFrom != nil && errTo != nil {
		return accountInfoList, ErrTransRepoDBWhileFetchingAccountInfo
	}

	//---------------Calculations
	cardFromInfo.Balance, errFrom = input.ProcessTransactionMinus(cardFromInfo.Balance, cardFromInfo.AccountRuleInfo)
	if errFrom != nil {
		return accountInfoList, errFrom
	}

	cardToInfo.Balance, errTo = input.ProcessTransactionPlus(cardToInfo.Balance, cardToInfo.AccountRuleInfo)
	if errTo != nil {
		return accountInfoList, errTo
	}

	//---------------updating
	resultFrom := accountRepoHandle.UpdateBalance(ctx, cardFromInfo.AccountID, cardFromInfo.Balance)
	resultTo := accountRepoHandle.UpdateBalance(ctx, cardToInfo.AccountID, cardToInfo.Balance)
	if resultFrom != nil || resultTo != nil {
		return accountInfoList, ErrTransRepoDBUpdatingBalanceFailed
	}

	err = t.InsertTransaction(ctx, cardFromInfo, cardToInfo, input.Amount())
	if err != nil {
		return accountInfoList, ErrTransRepoDBWhileInsertingTransaction
	}

	//---------------Commit Transaction
	if err = tx.Commit(); err != nil {
		//return fail(err)
	}
	//Make Sure Everything is Fine
	accountInfoList[0] = cardFromInfo
	accountInfoList[1] = cardFromInfo
	return accountInfoList, nil
}

func (t *TransactionRepositoryMySqlDB) InsertTransaction(ctx context.Context, fromAccount domain.AccountInfoOutput, toAccount domain.AccountInfoOutput, val int64) error {

	var (
		query string = fmt.Sprintf("INSERT INTO transaction (card_id_from, card_id_to, amount, transaction_type) VALUES (?, ?, ?, %d)", domain.TransactionTransfer)
	)

	tx := ctx.Value("tx").(*sql.Tx)
	if tx == nil {
		return ErrTransRepoDBInvalidTransactionExp
	}
	_, err := tx.ExecContext(ctx, query, fromAccount.CardId, toAccount.CardId, val)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionRepositoryMySqlDB(clientIn *sql.DB) *TransactionRepositoryMySqlDB {
	return &TransactionRepositoryMySqlDB{clientIn}
}
