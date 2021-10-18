package service

import (
	domain "github.com/io-m/bankapp/src/domain/customer"
)

type CustomerServicer interface {
	// Save(c domain.Customer) (string, error)
	GetAll() ([]*domain.CustomerResponse, error)
	GetOne(id int) (*domain.CustomerResponse, error)
}