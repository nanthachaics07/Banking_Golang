package service

import (
	"bank_test01/errs"
	logs "bank_test01/log"
	"bank_test01/repository"
	"database/sql"
	"log"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		logs.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	custResponses := []CustomerResponse{}

	for _, customer := range customers {
		var custResponse = CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")

		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError("Unexpected database error")

	}
	var custResponse = CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil

}
