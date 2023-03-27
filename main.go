package main

import (
	"fmt"
	"github.com/omid-h70/bank-service/cmd"
	"github.com/omid-h70/bank-service/internal/adapters/repository"
)

func main() {
	cmd.NewAppConfig().
		ServerAddress("localhost", "8000").
		CustomerRepo(repository.NewCustomerRepositoryMySqlDB()).
		NotifyService(repository.NewKaveNegarNotifyMsg()).
		Run()

	fmt.Println("Hi, i'm up")
}
