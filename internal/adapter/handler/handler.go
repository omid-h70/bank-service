package handler

import (
	"github.com/gorilla/mux"
	"github.com/omid-h70/bank-service/internal/adapter/response"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/pkg/errors"
	"net/http"
)

type AppHandler struct {
	accountHandler AccountHandler
	customerHandle CustomerHandler
	notifyHandler  domain.PushNotificationService
}

func (app *AppHandler) SetAppHandlers(router *mux.Router) {

	api := router.PathPrefix("/v1").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(defaultHandler)
	api.HandleFunc("/transfer", app.accountHandler.handleTransferCallBack).Methods(http.MethodPost)
	api.HandleFunc("/report", app.customerHandle.handleGetCustomerReport).Methods(http.MethodGet)
	api.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
}

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	response.NewError(errors.New("Invalid request"), http.StatusBadRequest).Send(w)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	response.NewSuccess("Yo I'm up", http.StatusOK).Send(w)
}

func (app *AppHandler) RegisterService(customer domain.CustomerService, transaction domain.TransactionService, notify domain.PushNotificationService) {
	app.customerHandle.service = customer
	app.accountHandler.service = transaction
	app.notifyHandler = notify
	app.accountHandler.notifyService = notify
}

func NewAppHandler() AppHandler {
	return AppHandler{}
}
