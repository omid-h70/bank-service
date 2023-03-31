package domain

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrInvalidMinTransactionAmount = errors.New("invalid Min Transaction Amount")
	ErrInvalidMaxTransactionAmount = errors.New("invalid Max Transaction Amount")
	ErrInvalidAccountBalance       = errors.New("invalid Account Balance")
)

type (
	TransactionService interface {
		ExecuteCardTransfer(context.Context, Transaction) ([]AccountInfoOutput, error)
	}

	Transaction struct {
		id       string
		cardFrom Card
		cardTo   Card
		amount   int64
	}

	TransactionRepository interface {
		MakeTransferFromCardToCard(ctx context.Context, input Transaction) ([]AccountInfoOutput, error)
		InsertTransaction(ctx context.Context, fromAccount AccountInfoOutput, toAccount AccountInfoOutput, val int64) error
	}
)

func (t *Transaction) Amount() int64 {
	return t.amount
}

func (t *Transaction) SetAmount(transactionAmount int64) {
	t.amount = transactionAmount
}

func (t *Transaction) CardTo() Card {
	return t.cardTo
}

func (t *Transaction) SetCardToInfo(card Card) {
	t.cardTo = card
}

func (t *Transaction) CardFrom() Card {
	return t.cardFrom
}

func (t *Transaction) SetCardFromInfo(card Card) {
	t.cardFrom = card
}

func (t *Transaction) ProcessTransactionMinus(amount int64, rule AccountRule) (int64, error) {
	tmpAmount := amount
	err := t.isTransactionValid(amount, rule)
	if err != nil {
		return tmpAmount, err
	}

	tmpAmount = amount - t.amount - rule.TransactionFee
	if tmpAmount < 0 {
		return amount, ErrInvalidAccountBalance
	}

	return tmpAmount, nil
}

func (t *Transaction) ProcessTransactionPlus(amount int64, rule AccountRule) (int64, error) {
	tmpAmount := amount
	err := t.isTransactionValid(amount, rule)
	if err != nil {
		return tmpAmount, err
	}

	tmpAmount = amount + t.amount
	//TODO ::: no rules defined here

	return tmpAmount, nil
}

func (t *Transaction) isTransactionValid(amount int64, rule AccountRule) error {
	if amount < rule.MinAmount {
		return ErrInvalidMinTransactionAmount
	}

	if amount > rule.MaxAmount {
		return ErrInvalidMaxTransactionAmount
	}

	return nil
}

func NewTransaction() Transaction {
	return Transaction{}
}
