package main

import (
	"fmt"
	"github.com/omid-h70/bank-service/cmd"
	"github.com/omid-h70/bank-service/internal/adapter/action"
	"github.com/omid-h70/bank-service/internal/adapter/handler"
	"github.com/omid-h70/bank-service/internal/adapter/repository"
	"github.com/omid-h70/bank-service/internal/core/service"
	"os"
)

func main() {

	appTransferAction := action.NewCreateTransferAction(repository.NewAccountRepositoryMySqlDB(os.Getenv("MYSQL_DB")))

	//appTransferAction := action.NewCreateTransferAction(repository.NewCustomerRepositoryMock())
	appTransferService := service.NewTransferService(appTransferAction, 1000)
	appHandler := handler.NewAccountHandler(appTransferService)

	cmd.NewAppConfig().
		ServerAddress("localhost", "8000").
		AppHandler(appHandler).
		NotifyService(repository.NewKaveNegarNotifyMsg()).
		Run()

	fmt.Println("Hi, i'm up")
}
