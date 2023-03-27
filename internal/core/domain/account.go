package domain

type Account struct {
	Id         string `json:"account_id"`
	CustomerId string `json:"customer_id"`
	AccountNum string `json:"account_number"`
}

type AccountRule struct {
	Id             string `json:"account_rule_id"`
	MinAmount      string
	MaxAmount      string
	TransactionFee string
}
