package usecase

import (
	"errors"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
)

// SearchBeerById list the detail of the beer brand
func SearchBeerById(br domain.BeerRepository, beerId int64) (dto.BeerItemResponse, error) {
	beer, err := br.Get(beerId)
	if err != nil {
		return dto.BeerItemResponse{}, err
	}
	if beer == nil {
		return dto.BeerItemResponse{}, errors.New("el Id de la cerveza no existe")
	}
	return dto.BeerItemResponse{
		ID:       beer.ID,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		Country:  beer.Country,
		Price:    beer.Price,
		Currency: beer.Currency,
	}, nil
}
