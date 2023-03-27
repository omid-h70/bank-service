package domain

type Transaction struct {
	Id     string `json:"transaction_id"`
	CardId string `json:"card_id"`
}
