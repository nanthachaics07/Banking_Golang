package handler

import (
	"bank_test01/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) CustomerHandler {
	return CustomerHandler{custSrv: custSrv}
}

func (h CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
