package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omid-h70/bank-service/internal/adapter/handler"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerConfig struct {
	Addr string
	Port string
}

type AppConfig struct {
	serverCnf  ServerConfig
	appHandler handler.AppHandler
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		appHandler: handler.NewAppHandler(),
	}
}

func (cnf *AppConfig) RegisterService(customer service.CustomerService, transaction service.TransactionService, notify domain.PushNotificationService) *AppConfig {
	cnf.appHandler.RegisterService(customer, transaction, notify)
	return cnf
}

//func (cnf *AppConfig) CustomerRepo(repo domain.CustomerRepository) *AppConfig {
//	cnf.repo = repo
//	return cnf
//}
//
//func (cnf *AppConfig) NotifyService(notifyRepo domain.PushNotificationService) *AppConfig {
//	cnf.notifyRepo = notifyRepo
//	return cnf
//}

func (cnf *AppConfig) ServerAddress(servConfig ServerConfig) *AppConfig {
	cnf.serverCnf = servConfig
	return cnf
}

func (cnf *AppConfig) Run() {
	router := mux.NewRouter()

	cnf.appHandler.SetAppHandlers(router)
	fmt.Println("Try to Run Server On " + cnf.serverCnf.Addr + ":" + cnf.serverCnf.Port)
	cnf.listen(router)
}

func (cnf *AppConfig) listen(router *mux.Router) {

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf("%s:%s", cnf.serverCnf.Addr, cnf.serverCnf.Port),
		Handler:      router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed")
	}

	log.Fatal("Service down")
}
