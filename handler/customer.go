package handler

import (
	"bank_test01/service"
	"net/http"
)

type CustomerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) CustomerHandler {
	return CustomerHandler{custSrv: custSrv}
}

func (h CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	println("Hello World")
}
