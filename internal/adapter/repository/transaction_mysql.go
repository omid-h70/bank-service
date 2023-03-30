package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/pkg/errors"
)

var (
	ErrTransRepoDBInvalidTransactionExp = errors.New("invalid NULL Transaction Exception")
	ErrTransRepoDBUpdatingBalanceFailed = errors.New("Error while updating Balance")
)

type TransactionRepositoryMySqlDB struct {
	client *sql.DB
}

func (t *TransactionRepositoryMySqlDB) MakeTransferFromCardToCard(ctx context.Context, input domain.Transaction) error {

	var (
		cardFrom = input.CardFrom()
		cardTo   = input.CardTo()
		err      error
	)

	tx, err := t.client.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "error updating account balance")
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
		return ErrTransRepoDBUpdatingBalanceFailed
	}

	//---------------Calculations
	cardFromInfo.Balance, errFrom = input.ProcessTransactionMinus(cardFromInfo.Balance, cardFromInfo.AccountRuleInfo)
	if errFrom != nil {
		return errFrom
	}

	cardToInfo.Balance, errTo = input.ProcessTransactionPlus(cardToInfo.Balance, cardToInfo.AccountRuleInfo)
	if errTo != nil {
		return errTo
	}

	//---------------updating
	resultFrom := accountRepoHandle.UpdateBalance(ctx, cardFromInfo.AccountID, cardFromInfo.Balance)
	resultTo := accountRepoHandle.UpdateBalance(ctx, cardToInfo.AccountID, cardToInfo.Balance)
	if resultFrom != nil || resultTo != nil {
		return errors.Wrap(err, "error updating account balance")
	}

	err = t.InsertTransaction(ctx, cardFromInfo, cardToInfo, input.Amount())
	if err != nil {
		return err
	}

	//---------------Commit Transaction
	if err = tx.Commit(); err != nil {
		//return fail(err)
	}

	return nil
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
