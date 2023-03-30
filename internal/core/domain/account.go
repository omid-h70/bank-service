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

type (
	// TransferService input port

	// CreateTransferPresenter output port
	//CreateTransferPresenter interface {
	//	Output(domain.Transfer) CreateTransferOutput
	//}

	// CardTransferOutput output data

	AccountInfoOutput struct {
		AccountID       string `json:"account_id"`
		CardId          string `json:"card_id"`
		Balance         int64  `json:"card_balance"`
		TransferTime    string `json:"date_time"`
		AccountRuleInfo AccountRule
		PhoneNumber     string `json:"phone_num"`
	}
)

const (
	TransactionTransfer int = iota
)

const (
	AtomicInfo int = iota
	NormalInfo
)

type (
	AccountRepository interface {
		GetAccountInfoByCard(ctx context.Context, cardNum string, mode int) (AccountInfoOutput, error)
		UpdateAccountBalance(ctx context.Context, accountId string, val int64) error
	}
)
