package usecase

import (
	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
)

// AddBeers enter a new beer
func AddBeer(br domain.BeerRepository, beerItem dto.BeerItemRequest) error {
	return br.Save(beerItem.Name, beerItem.Brewery, beerItem.Country, beerItem.Price, beerItem.Currency)
}
