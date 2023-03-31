package handler

import (
	"github.com/omid-h70/bank-service/internal/adapter/response"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"net/http"
)

type CustomerHandler struct {
	service domain.CustomerService
}

func (c *CustomerHandler) handleGetCustomerReport(w http.ResponseWriter, r *http.Request) {
	//it sends to Response Page To Write

	customerList, err := c.service.GetMostActiveCustomersWithinTime(r.Context(), 10, 1000)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	response.NewSuccess(customerList, 200).Send(w)
}

func NewCustomerHandler(customerService domain.CustomerService) CustomerHandler {
	return CustomerHandler{service: customerService}
}
