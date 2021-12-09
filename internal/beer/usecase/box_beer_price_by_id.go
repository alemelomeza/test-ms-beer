package usecase

import (
	"strconv"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
)

// BoxBeerPriceById list the price of a case of brand beers
func BoxBeerPriceById(br domain.BeerRepository, cr domain.CurrencyRepository, beerBox dto.BeerBoxRequest) (dto.BeerBoxResponse, error) {
	beerID, err := strconv.ParseInt(beerBox.BeerID, 10, 64)
	if err != nil {
		return dto.BeerBoxResponse{}, err
	}
	if beerBox.Quantity == "" {
		beerBox.Quantity = "6"
	}
	quantity, err := strconv.ParseFloat(beerBox.Quantity, 64)
	if err != nil {
		return dto.BeerBoxResponse{}, err
	}
	beer, err := br.Get(beerID)
	if err != nil {
		return dto.BeerBoxResponse{}, err
	}
	if beer.Currency == beerBox.Currency {
		return dto.BeerBoxResponse{
			PriceTotal: int64(beer.Price * quantity),
		}, nil
	}
	beerCurrency, beerBoxCurrency, err := cr.Get(beer.Currency, beerBox.Currency)
	if err != nil {
		return dto.BeerBoxResponse{}, err
	}

	c := (beer.Price / beerCurrency) * beerBoxCurrency

	return dto.BeerBoxResponse{
		PriceTotal: int64(c * quantity),
	}, nil
}
