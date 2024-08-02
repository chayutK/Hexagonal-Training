package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	q := "select customer_id,name,date_of_birth,city,zipcode,status from customers"
	err := r.db.Select(&customers, q)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	q := "select customer_id,name,date_of_birth,city,zipcode,status from customers WHERE customer_id=?"
	err := r.db.Get(&customer, q, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
