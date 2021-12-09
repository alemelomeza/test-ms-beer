package dto

type BeerItemResponse struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type BeerBoxResponse struct {
	PriceTotal int64 `json:"priceTotal"`
}
