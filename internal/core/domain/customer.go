package domain

type Customer struct {
	Id       string `json:"customer_id"`
	PhoneNum string `json:"phone_num"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	//GetTopCustomers(options ...string)
}

type CustomerService interface {
	GetAllCustomers() ([]Customer, error)
}
