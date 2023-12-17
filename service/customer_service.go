package service

import "bank_test01/repository"

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	return nil, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	return nil, nil
}
