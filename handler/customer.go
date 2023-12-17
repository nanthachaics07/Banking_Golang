package handler

import "bank_test01/service"

type CustomerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) CustomerHandler {
	return CustomerHandler{custSrv: custSrv}
}
