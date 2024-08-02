package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", Zipcode: "10020", DateOfBirth: "2001-11-30", Status: 1},
		{CustomerID: 1001, Name: "Rob", City: "New Delhi", Zipcode: "10020", DateOfBirth: "2001-11-30", Status: 1},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}
func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
