package main

import (
	"github.com/io-m/bankapp/internal"
	"github.com/io-m/bankapp/internal/db"
	"github.com/io-m/bankapp/internal/logger"
	"github.com/io-m/bankapp/src/api"
	domain "github.com/io-m/bankapp/src/domain/customer/adapters"
	handler "github.com/io-m/bankapp/src/handler/customer"
	service "github.com/io-m/bankapp/src/service/customer"
)

func main() {
	config := internal.NewConfig()
	router := api.NewGorillaRouter()
	mysqlClient := db.NewDBClient(config.Driver, config.DBConn)
	handler := handler.NewCustomerHandler(service.NewCustomerService(domain.NewCustomerMySql(mysqlClient)))
	router.Get("/customers", handler.GetAll)
	router.Get("/customers/{id:[0-9 A-Z a-z]+}", handler.GetOne)
	if err := router.Serve(config.Addr, config.Port); err != nil {
		logger.Warn.Fatalf("Problem occured while listening on the port %s ", config.Port)
	}
}