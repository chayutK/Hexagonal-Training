package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"errors"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) customerService {
	return customerService{customerRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	resp := []CustomerResponse{}
	for _, customer := range customers {
		custResp := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		resp = append(resp, custResp)
	}
	return resp, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepo.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	resp := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &resp, nil
}
