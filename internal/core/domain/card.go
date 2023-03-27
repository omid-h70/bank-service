package domain

type Card struct {
	Id         string `json:"card_id"`
	AccountId  string `json:"account_id"`
	CardNumber string `json:"card_number"`
}
