package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omid-h70/bank-service/internal/adapters/handler"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type serverConfig struct {
	addr string
	port string
}

type AppConfig struct {
	serverCnf serverConfig
	//TODO: db repo
	repo domain.CustomerRepository
	//
	notifyRepo domain.PushNotificationService
}

func NewAppConfig() AppConfig {
	return AppConfig{}
}

func (cnf AppConfig) CustomerRepo(repo domain.CustomerRepository) AppConfig {
	cnf.repo = repo
	return cnf
}

func (cnf AppConfig) NotifyService(notifyRepo domain.PushNotificationService) AppConfig {
	cnf.notifyRepo = notifyRepo
	return cnf
}

func (cnf AppConfig) ServerAddress(addr string, port string) AppConfig {
	cnf.serverCnf = serverConfig{
		addr,
		port,
	}
	return cnf
}

func (cnf AppConfig) Run() {
	router := mux.NewRouter()

	//wiring
	ch := handler.CustomerHandler{services.NewCustomerService(cnf.repo)}
	//ch := handler.CustomerHandler{services.NewCustomerService(repository.NewCustomerRepositoryMySqlDB())}

	//define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	fmt.Println("Try to Run Server On " + cnf.serverCnf.addr + ":" + cnf.serverCnf.port)
	log.Fatal(http.ListenAndServe(cnf.serverCnf.addr+":"+cnf.serverCnf.port, router))
}

//func (cnf AppConfig) setAppHandlers(router *mux.Router) {
//	api := router.PathPrefix("/v1").Subrouter()
//
//	api.Handle("/transfers", g.buildCreateTransferAction()).Methods(http.MethodPost)
//	//api.Handle("/transfers", g.buildFindAllTransferAction()).Methods(http.MethodGet)
//	//
//	//api.Handle("/accounts/{account_id}/balance", g.buildFindBalanceAccountAction()).Methods(http.MethodGet)
//	//api.Handle("/accounts", g.buildCreateAccountAction()).Methods(http.MethodPost)
//	//api.Handle("/accounts", g.buildFindAllAccountAction()).Methods(http.MethodGet)
//
//	api.HandleFunc("/health", action.HealthCheck).Methods(http.MethodGet)
//}

func (cnf AppConfig) listen(router *mux.Router) {
	//g.setAppHandlers(g.router)
	//g.middleware.UseHandler(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		//Addr:         fmt.Sprintf(":%d", g.port),
		//Handler:      g.middleware,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		//g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			//g.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		//g.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	//g.log.Infof("Service down")
}
