package main

import (
	"bank_test01/handler"
	"bank_test01/repository"
	"bank_test01/service"
	"fmt"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	initConfig()

	db, err := sqlx.Open("mysql", "root:babana@(localhost:3306)/banking?parseTime=true")
	if err != nil {
		panic(err)
	}

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryDB

	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	go func() {
		fmt.Println("Server is listening on port 8000")
		if err := http.ListenAndServe(":8000", router); err != nil {
			panic(err)
		}
	}()
	// Block the main goroutine to keep the server running
	select {}

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}
