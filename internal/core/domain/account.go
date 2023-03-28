package domain

import (
	"context"
)

type Account struct {
	Id             string `json:"account_id"`
	AccountRuleId  string `json:"account_rule_id"`
	CustomerId     string `json:"customer_id"`
	AccountNum     string `json:"account_number"`
	AccountBalance string `json:"account_balance"`
}

type AccountRule struct {
	Id             string `json:"account_rule_id"`
	MinAmount      string `json:"min_allowed_amount"`
	MaxAmount      string `json:"max_allowed_amount"`
	TransactionFee string `json:"transaction_fee"`
}

type (
	// TransferService input port
	TransferService interface {
		ExecuteCardTransfer(context.Context, CardTransferInput) (CardTransferOutput, error)
	}

	// CardInfo input data
	CardInfo struct {
		AccountId string `json:"account_id" validate:"required,uuid4"`
		CardNum   string `json:"account_from_id" `
	}

	CardTransferInput struct {
		CardFrom CardInfo
		CardTo   CardInfo
		Amount   int64 `json:"amount" validate:"gt=0,required"`
	}

	// CreateTransferPresenter output port
	//CreateTransferPresenter interface {
	//	Output(domain.Transfer) CreateTransferOutput
	//}

	// CardTransferOutput output data
	CardTransferOutput struct {
		ID           string `json:"id"`
		CardFrom     CardInfo
		CardTo       CardInfo
		Amount       float64 `json:"amount"`
		TransferTime string  `json:"date_time"`
	}
)

const (
	TransactionTransfer int = iota
)

const (
	AtomicBalance int = iota
	NormalBalance
)

type (
	AccountRepositoryDB interface {
		InsertTransaction(ctx context.Context, input CardTransferInput) error
		GetBalanceByCard(ctx context.Context, accountId string, mode int) (int64, error)
		GetBalanceByAccount(ctx context.Context, accountId string, mode int) (int64, error)
		UpdateBalance(ctx context.Context, accountId string, val int64) error
		MakeTransferFromCardToCard(ctx context.Context, input CardTransferInput) (CardTransferOutput, error)
	}
)
