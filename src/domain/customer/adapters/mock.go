package domain

import domain "github.com/io-m/bankapp/src/domain/customer"

type customerMock struct {
	customers []domain.Customer
}

func NewCustomerMock() domain.CustomerRepo {
	customers := []domain.Customer{
		{
			Name:    "Josip",
			Age:     29,
			Address: "Hannemanns Alle 4a",
			Country: "Denmark",
			Active:  true,
		},
		{
			Name:    "Martina",
			Age:     29,
			Address: "Hannemanns Alle 4a",
			Country: "Denmark",
			Active:  false,
		},
	}
	return &customerMock{
		customers: customers,
	}
}

func (cm *customerMock) Save(c domain.Customer) (string, error) {
	return "", nil
}
func (cm *customerMock) GetAll() ([]domain.Customer, error) {
	return cm.customers, nil
}
func (cm *customerMock) GetOne(id int) (*domain.Customer, error) {
	return nil, nil
}
