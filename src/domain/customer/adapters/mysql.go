package domain

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/io-m/bankapp/internal/logger"
	domain "github.com/io-m/bankapp/src/domain/customer"
)
const (
	SQL_GET_ALL = `SELECT * FROM bankaccount.User;`
	SQL_GET_ONE = `SELECT * FROM bankaccount.User
	WHERE id = ?;
	`
)


type customerMySql struct {
	client *sql.DB
}

func NewCustomerMySql(cl *sql.DB) domain.CustomerRepo {
	return &customerMySql {
		client: cl,
	}
}

func (ms *customerMySql) GetAll() ([]domain.Customer, error) {
	rows, err := ms.client.Query(SQL_GET_ALL)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No customers")
		}
		logger.Warn.Println(err.Error())
		return nil, err
	}
	customers := []domain.Customer{}
	for rows.Next() {
		var c domain.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Age, &c.Address, &c.Country, &c.Active); err != nil {
			logger.Warn.Fatal("Could not scan into customer")
		}
		customers = append(customers, c)
	}
	return customers, nil
}
func (ms *customerMySql) GetOne(id int) (*domain.Customer, error) {
	customer := domain.Customer{}
	if err := ms.client.QueryRow(SQL_GET_ONE, id).Scan(&customer.ID, &customer.Name, &customer.Age, &customer.Address, &customer.Country, &customer.Active); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Could not find a selected customer")
		}
		logger.Warn.Println("Could not retrieve that customer")
		return nil, err
	}
	return &customer, nil
}