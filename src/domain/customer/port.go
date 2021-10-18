package domain

type CustomerRepo interface {
	// Save(c Customer) (string, error)
	GetAll() ([]Customer, error)
	GetOne(id int) (*Customer, error)
}