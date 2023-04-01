package handler

import (
	"github.com/omid-h70/bank-service/internal/adapter/response"
	"github.com/omid-h70/bank-service/internal/core/service"
	"net/http"
	"strconv"
)

const (
	DefaultReportMinutesTime = 10
)

type CustomerHandler struct {
	service service.CustomerService
}

func (c *CustomerHandler) handleGetCustomerReport(w http.ResponseWriter, r *http.Request) {
	//it sends to Response Page To Write
	values := r.URL.Query()
	time := values.Get("time")
	minutes, _ := strconv.ParseInt(time, 10, 0)
	if minutes == 0 {
		minutes = DefaultReportMinutesTime
	}

	customerList, err := c.service.GetMostActiveCustomersWithinTime(r.Context(), 10, int(minutes))
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	} else if len(customerList) == 0 {
		response.NewSuccess("No Result Within Given Time", 200).Send(w)
		return
	}
	response.NewSuccess(customerList, 200).Send(w)
}

func NewCustomerHandler(customerService service.CustomerService) CustomerHandler {
	return CustomerHandler{service: customerService}
}
