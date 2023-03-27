package repository

import "github.com/omid-h70/bank-service/internal/core/domain"

type CustomerRepositoryMock struct {
	customers []domain.Customer
}

func (s CustomerRepositoryMock) FindAll() (customer []domain.Customer, err error) {
	//customers := []Customer{
	//	{"omid", "989123993699"},
	//	{"omid", "989123993699"},
	//}
	return s.customers, nil
}

func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []domain.Customer{
		{"omid", "989123993699"},
		{"omid", "989123993699"},
	}
	return CustomerRepositoryMock{customers}
}
