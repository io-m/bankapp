package domain

type Customer struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Age     int    `db:"age"`
	Address string `db:"address"`
	Country string `db:"country"`
	Active  bool   `db:"active"`
}

// ToDTO converts domain object to data transfer object
func (c Customer) ToDTO() *CustomerResponse {
	return &CustomerResponse{
		ID:      c.ID,
		Name:    c.Name,
		Age:     c.Age,
		Address: c.Address,
		Country: c.Country,
		Active:  c.Active,
	}
}