package repository

import "errors"

type CustomerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "John Doe", DateOfBirth: "1990-01-01", City: "New York", ZipCode: "10001", Status: 1},
		{CustomerID: 1002, Name: "Jane Smith", DateOfBirth: "1991-02-02", City: "Los Angeles", ZipCode: "90001", Status: 1},
	}

	return CustomerRepositoryMock{customers}
}

func (r CustomerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r CustomerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
