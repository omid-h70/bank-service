package handler

import (
	"encoding/json"
	"fmt"
	"github.com/omid-h70/bank-service/internal/adapter/response"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/service"
	"github.com/pkg/errors"
	"net/http"
)

var (
	errInvalidApplicationType = errors.New("Request Application Type Must be jsob")
)

type AccountHandler struct {
	service service.TransferTransaction
}

func NewAccountHandler(service service.TransferTransaction) AccountHandler {
	return AccountHandler{service}
}

func (a *AccountHandler) HandleTransferCallBack(w http.ResponseWriter, r *http.Request) {

	fmt.Println("i`m in 1")
	var input domain.CardTransferInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		//TODO Log your error
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		response.NewError(errInvalidApplicationType, http.StatusBadRequest).Send(w)
		return
	}

	//it sends to Response Page To Write
	fmt.Println(input)

	output, err := a.service.ExecuteCardTransfer(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	fmt.Println(output)
	response.NewSuccess("done", 200)
}
