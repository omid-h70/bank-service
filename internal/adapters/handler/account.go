package handler

import (
	"github.com/omid-h70/bank-service/internal/core/services"
	"net/http"
)

type AccountHandler struct {
	service services.TransferTransaction
}

func (a *AccountHandler) HandleTransferCallBack(w http.ResponseWriter, r *http.Request) {

	//var input services.CreateTransferInput
	//if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	//	//logging.NewError(
	//	//	t.log,
	//	//	err,
	//	//	t.logKey,
	//	//	http.StatusBadRequest,
	//	//).Log(t.logMsg)
	//
	//	response.NewError(err, http.StatusBadRequest).Send(w)
	//	return
	//}
	//defer r.Body.Close()
	//it sends to Response Page To Write

	//services.CreateTransferInput{
	//
	//}
	//TODO: add input and make transaction
	//customerList, err := c.service.Execute()
	//if err != nil {
	//	log.Fatal("Err1")
	//}
	//
	//if r.Header.Get("Content-Type") == "application/json" {
	//	w.Header().Add("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(customerList)
	//}
}
