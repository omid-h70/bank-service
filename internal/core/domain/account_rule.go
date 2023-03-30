package domain

type (
	AccountRule struct {
		Id             string `json:"account_rule_id"`
		MinAmount      int64  `json:"min_allowed_amount"`
		MaxAmount      int64  `json:"max_allowed_amount"`
		TransactionFee int64  `json:"transaction_fee"`
	}
)
