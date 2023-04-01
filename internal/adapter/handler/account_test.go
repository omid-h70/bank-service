package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	TRANSFER_URL   = "/transfer"
	CARD_FROM_TEST = ""
	CARD_TO_TEST   = ""
)

type accountInterfaceMock struct {
}

type pushNotificationMock struct {
}

func (pushNotificationMock) GetReceiverNotifyMessage(_ domain.AccountInfoOutput, _ string) string {
	return ""
}

func (pushNotificationMock) GetSenderNotifyMessage(_ domain.AccountInfoOutput, _ string) string {
	return ""
}

func (pushNotificationMock) SendNotifyMessage(sender string, receptor []string, msg string) error {
	return nil
}

func (a accountInterfaceMock) ExecuteCardTransfer(_ context.Context, t domain.Transaction) ([]domain.AccountInfoOutput, error) {
	cardFrom := t.CardFrom()
	cardTo := t.CardTo()

	mockData := []domain.AccountInfoOutput{
		{
			AccountID:    "1001",
			CardId:       cardFrom.CardNum(),
			Balance:      1000,
			TransferTime: fmt.Sprintf("%s", time.Now()),
			AccountRuleInfo: domain.AccountRule{
				MinAmount:      1000,
				MaxAmount:      50000000,
				TransactionFee: 500,
			},
			CustomerInfo: domain.Customer{
				PhoneNum: "+989123993699",
			},
		},
		{
			AccountID:    "1002",
			CardId:       cardTo.CardNum(),
			Balance:      1000,
			TransferTime: fmt.Sprintf("%s", time.Now()),
			AccountRuleInfo: domain.AccountRule{
				MinAmount:      1000,
				MaxAmount:      50000000,
				TransactionFee: 500,
			},
			CustomerInfo: domain.Customer{
				PhoneNum: "+903934262",
			},
		},
	}
	return mockData, nil
}

func Test_transfer_should_return_fail_when_card_number_len_is_not_valid(t *testing.T) {
	var jsonData = []byte(`{
		"card_from_number" : "195720212030",
		"card_to_number": "622106106090",
		"transaction_amount":"123456"
	}`)

	router := mux.NewRouter()

	accountHandler := NewAccountHandler(accountInterfaceMock{}, pushNotificationMock{})
	router.HandleFunc(TRANSFER_URL, accountHandler.handleTransferCallBack).Methods(http.MethodPost)

	request, _ := http.NewRequest(http.MethodPost, TRANSFER_URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Error("Test Failed")
	}
}

func Test_transfer_should_return_fail_if_card_number_pattern_is_not_valid(t *testing.T) {
	var jsonData = []byte(`{
		"card_from_number" : "1234123412341234",
		"card_to_number": "4567456745674567",
		"transaction_amount":"123456"
	}`)

	router := mux.NewRouter()

	accountHandler := NewAccountHandler(accountInterfaceMock{}, pushNotificationMock{})
	router.HandleFunc(TRANSFER_URL, accountHandler.handleTransferCallBack).Methods(http.MethodPost)

	request, _ := http.NewRequest(http.MethodPost, TRANSFER_URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Error("Test Failed")
	}
}
