package main

import (
	"bank_test01/repository"
	"bank_test01/service"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:babana@(localhost:3306)/banking?parseTime=true")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)

	customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)
}
