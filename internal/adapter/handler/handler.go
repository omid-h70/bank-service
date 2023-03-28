package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/omid-h70/bank-service/internal/adapter/action"
	"github.com/omid-h70/bank-service/internal/core/service"

	"net/http"
)

type AppHandler struct {
	accountHandler AccountHandler
}

type CustomerHandler struct {
	Service service.CustomerService
}

func (app AppHandler) AccountHandler(hndl AccountHandler) AppHandler {
	app.accountHandler = hndl
	return app
}

func (app AppHandler) SetAppHandlers(router *mux.Router) {

	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/transfer", app.accountHandler.HandleTransferCallBack).Methods(http.MethodPost)
	//api.Handle("/transfers", g.buildFindAllTransferAction()).Methods(http.MethodGet)
	//
	//api.Handle("/accounts/{account_id}/balance", g.buildFindBalanceAccountAction()).Methods(http.MethodGet)
	//api.Handle("/accounts", g.buildCreateAccountAction()).Methods(http.MethodPost)
	//api.Handle("/accounts", g.buildFindAllAccountAction()).Methods(http.MethodGet)

	api.HandleFunc("/health", action.HealthCheck).Methods(http.MethodGet)
}

func (c *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
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
