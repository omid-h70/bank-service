package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/omid-h70/bank-service/internal/adapter/response"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

var (
	ErrInvalidApplicationType = errors.New("Request Application Type Must be json")
)

type AccountHandler struct {
	service domain.TransactionService
}

type TransactionRequest struct {
	CardFromNum       string `json:"card_from_number" validate:"required"`
	CardToNum         string `json:"card_to_number" validate:"required"`
	TransactionAmount string `json:"transaction_amount" validate:"required"`
}

func (a *AccountHandler) handleTransferCallBack(w http.ResponseWriter, r *http.Request) {

	var req TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		response.NewError(ErrInvalidApplicationType, http.StatusBadRequest).Send(w)
		return
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	//Mapping Request To Domain

	inAmount, parseErr := strconv.ParseInt(req.TransactionAmount, 10, 64)
	if parseErr != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	input := domain.NewTransaction()
	input.SetCardFromInfo(domain.NewCard(req.CardFromNum))
	input.SetCardToInfo(domain.NewCard(req.CardToNum))
	input.SetAmount(inAmount)

	err = a.service.ExecuteCardTransfer(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	//fmt.Println(output)
	response.NewSuccess("done", 200).Send(w)
}

func NewAccountHandler(service domain.TransactionService) AccountHandler {
	return AccountHandler{service: service}
}
