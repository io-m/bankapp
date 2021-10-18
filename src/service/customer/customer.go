package service

import (
	"github.com/io-m/bankapp/internal/logger"
	domain "github.com/io-m/bankapp/src/domain/customer"
)

type CustomerService struct {
	repo domain.CustomerRepo
}

func NewCustomerService(r domain.CustomerRepo) CustomerServicer {
	return &CustomerService{
		repo: r,
	}
}

func (cs *CustomerService) Save(c domain.Customer) (string, error) {
	return "", nil
}
func (cs *CustomerService) GetAll() ([]*domain.CustomerResponse, error) {
	customers, err := cs.repo.GetAll()
	if err != nil {
		logger.Warn.Println(err.Error())
		return nil, err
	}
	customersResponse := []*domain.CustomerResponse{}
	for _, c := range customers {
		customersResponse = append(customersResponse, c.ToDTO())
	}
	return customersResponse, nil
}
func (cs *CustomerService) GetOne(id int) (*domain.CustomerResponse, error) {
	c, err := cs.repo.GetOne(id)
	if err != nil {
		logger.Warn.Println(err.Error())
		return nil, err
	}
	customerResponse := c.ToDTO()
	return customerResponse, nil
}

