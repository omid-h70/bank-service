package handler

import (
	"encoding/json"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/service"
	"net/http"
)

//var (
//	ErrCInvalidApplicationType = errors.New("Request Application Type Must be jsob")
//)

type CustomerHandler struct {
	service domain.CustomerService
}

func (c *CustomerHandler) handleGetCustomerReport(w http.ResponseWriter, r *http.Request) {
	//it sends to Response Page To Write

	//ustomerList, err := c.Service.GetAllCustomers()
	//if err != nil {
	//	log.Fatal("Err1")
	//}
	//
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("customerList")
	}
}

func NewCustomerHandler(service service.ReportService) CustomerHandler {
	return CustomerHandler{service: service}
}
