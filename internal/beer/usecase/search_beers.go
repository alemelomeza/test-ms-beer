package usecase

import (
	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
)

// SearchBeers list all beers
func SearchBeers(br domain.BeerRepository) ([]dto.BeerItemResponse, error) {
	list, err := br.List()
	if err != nil {
		return nil, err
	}

	beers := make([]dto.BeerItemResponse, len(list))
	for i, beer := range list {
		beers[i] = dto.BeerItemResponse{
			ID:       beer.ID,
			Name:     beer.Name,
			Brewery:  beer.Brewery,
			Country:  beer.Country,
			Price:    beer.Price,
			Currency: beer.Currency,
		}
	}

	return beers, nil
}
