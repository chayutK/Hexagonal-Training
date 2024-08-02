package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, req NewAccountRequest) (*AccountResponse, error) {
	//NOTE - Validate Input
	if req.Amount < 5000 {
		return nil, errs.NewValidationError("amount must be at least 5000")
	}
	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-2 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}

	newAcc, err := s.accountRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	accResp := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &accResp, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accountRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	resp := []AccountResponse{}
	for _, account := range accounts {
		accResp := AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		}
		resp = append(resp, accResp)
	}

	return resp, nil
}
