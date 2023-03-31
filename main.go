package main

import (
	"fmt"
	"github.com/omid-h70/bank-service/cmd"
	"github.com/omid-h70/bank-service/internal/adapter/notification"
	"github.com/omid-h70/bank-service/internal/adapter/repository"
	"github.com/omid-h70/bank-service/internal/core/service"
	"path/filepath"
	"runtime"
)

func ProjectRootDir() string {
	_, d, _, _ := runtime.Caller(1)
	return filepath.Dir(d)
}

func main() {

	//Setting Database
	//appDbClient := repository.NewRepositoryMySqlDB(os.Getenv("MYSQL_DB"))
	appDbClient := repository.NewRepositoryMySqlDB("webservicedb")

	//Setting Services And Handlers
	appTransactionService := service.NewTransactionService(repository.NewTransactionRepositoryMySqlDB(appDbClient), 1000)
	appCustomerService := service.NewCustomerService(repository.NewCustomerRepositoryMySqlDB(appDbClient), 1000)
	notificationService := service.NewPushNotificationService(notification.NewKaveNegarNotifyMsg("1234"), 1000)

	cmd.NewAppConfig().
		ServerAddress("localhost", "8000").
		RegisterService(appCustomerService, appTransactionService, notificationService).
		Run()

	fmt.Println("Hi, i'm up")
}
