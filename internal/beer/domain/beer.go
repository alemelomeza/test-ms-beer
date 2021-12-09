package domain

type BeerRepository interface {
	List() ([]Beer, error)
	Get(ID int64) (*Beer, error)
	Save(name string, brewery string, country string, price float64, currency string) error
}

type CurrencyRepository interface {
	Get(currencyBeer string, currencyBeerBox string) (float64, float64, error)
}

type Beer struct {
	ID       int64
	Name     string
	Brewery  string
	Country  string
	Price    float64
	Currency string
}
