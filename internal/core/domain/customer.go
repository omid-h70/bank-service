package domain

import (
	"context"
)

type Customer struct {
	Id       string `json:"customer_id"`
	PhoneNum string `json:"phone_num"`
}

type CustomerReportOut struct {
	CustomerID      string
	TransactionId   string
	CardIdFrom      string
	CardIdTo        string
	Amount          int64
	TransactionType int
	TransactionTime string
	Index           int
}

type CustomerRepository interface {
	FindMostActiveCustomersWithinTime(ctx context.Context, count int, time int) ([]CustomerReportOut, error)
}
