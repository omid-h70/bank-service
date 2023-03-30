package handler

import (
	"github.com/gorilla/mux"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"net/http"
)

type AppHandler struct {
	accountHandler AccountHandler
	customerHandle CustomerHandler
	notifyHandler  domain.PushNotificationService
}

func (app *AppHandler) SetAppHandlers(router *mux.Router) {

	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/transfer", app.accountHandler.handleTransferCallBack).Methods(http.MethodPost)
	api.HandleFunc("/report", app.customerHandle.handleGetCustomerReport).Methods(http.MethodGet)
	api.HandleFunc("/health", healthCheck).Methods(http.MethodGet)

}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app *AppHandler) RegisterService(services ...any) {
	for _, element := range services {
		switch service := element.(type) {
		case domain.CustomerService:
			app.customerHandle.service = service
		case domain.TransactionService:
			app.accountHandler.service = service
		case domain.PushNotificationService:
			app.notifyHandler = service
		}
	}
}

func NewAppHandler() AppHandler {
	return AppHandler{}
}
