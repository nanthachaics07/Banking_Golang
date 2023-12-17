package main

import (
	"bank_test01/repository"
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
	_ = customerRepository

	// 	customers, err := customerRepository.GetAll()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(customers)
	// }

	customers, err := customerRepository.GetById(2000)
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}
