package dto

type BeerItemRequest struct {
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type BeerBoxRequest struct {
	BeerID   string `json:"beerID"`
	Currency string `json:"currency"`
	Quantity string `json:"quantity"`
}
