package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	q := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) values (?,?,?,?,?)"
	result, err := r.db.Exec(q, acc.CustomerID, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)
	return &acc, nil
}

func (r accountRepositoryDB) GetAll(id int) ([]Account, error) {
	accounts := []Account{}

	q := "SELECT account_id,customer_id,opening_date,account_type,amount,status from accounts where customer_id=?"
	err := r.db.Select(&accounts, q, id)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
