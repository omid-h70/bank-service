package main

import (
	"fmt"
	"github.com/omid-h70/bank-service/cmd"
	"github.com/omid-h70/bank-service/internal/adapter/notification"
	"github.com/omid-h70/bank-service/internal/adapter/repository"
	"github.com/omid-h70/bank-service/internal/core/service"
	"os"
	"path/filepath"
	"runtime"
)

func ProjectRootDir() string {
	_, d, _, _ := runtime.Caller(1)
	return filepath.Dir(d)
}

var (
	dbTestClientConfig = repository.MySqlConfig{
		DbServerAddr: "mysqldb",
		DbServerPort: "3306",
		DbName:       "webServiceDB",
		DbUser:       "root",
		DbPass:       "root",
	}
	testServerConfig = cmd.ServerConfig{
		Addr: "0.0.0.0",
		Port: "8000",
	}
)

func main() {

	//Setting Database
	dbClientConfig := repository.MySqlConfig{
		DbServerAddr: os.Getenv("MYSQL_CONTAINER_NAME"),
		DbServerPort: os.Getenv("MYSQL_CONTAINER_PORT"),
		DbName:       os.Getenv("MYSQL_DATABASE"),
		DbUser:       os.Getenv("MYSQL_USER"),
		DbPass:       os.Getenv("MYSQL_PASS"),
	}
	dbClientConfig = dbTestClientConfig
	fmt.Println("Db Config", dbClientConfig)
	appDbClient := repository.NewRepositoryMySqlDB(dbClientConfig)

	//Setting Services And Handlers
	appTransactionService := service.NewTransactionService(repository.NewTransactionRepositoryMySqlDB(appDbClient), 1000)
	appCustomerService := service.NewCustomerService(repository.NewCustomerRepositoryMySqlDB(appDbClient), 1000)
	notificationService := service.NewPushNotificationService(ProjectRootDir(), notification.NewKaveNegarNotifyMsg("1234"), 1000)

	serverConfig := cmd.ServerConfig{
		Addr: os.Getenv("APP_SERVER_ADDR"),
		Port: os.Getenv("APP_HOST_PORT"),
	}
	//serverConfig = testServerConfig

	cmd.NewAppConfig().
		ServerAddress(serverConfig).
		RegisterService(appCustomerService, appTransactionService, notificationService).
		Run()

	fmt.Println("Hi, i'm up")
}
