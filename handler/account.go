package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSvc service.AccountService
}

func NewAccountHandler(accSvc service.AccountService) accountHandler {
	return accountHandler{accSvc: accSvc}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	req := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	resp, err := h.accSvc.NewAccount(customerID, req)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	resp, err := h.accSvc.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
