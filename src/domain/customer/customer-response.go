package domain

type CustomerResponse struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Country string `json:"country,omitempty"`
	Active  bool   `json:"active,omitempty"`
}
