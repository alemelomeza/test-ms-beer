package usecase

import (
	"testing"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
	mocks_respository "github.com/alemelomeza/test-ms-beer/internal/beer/infra/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBoxBeerPriceById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks_respository.NewMockBeerRepository(ctrl)
	repositoryMock.EXPECT().Get(gomock.Any()).Return(&domain.Beer{
		ID:       1,
		Name:     "Golden",
		Brewery:  "Kross",
		Country:  "Chile",
		Price:    10.5,
		Currency: "EUR",
	}, nil)
	currencyMock := mocks_respository.NewMockCurrencyRepository(ctrl)
	currencyMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(0.883155, 839.497922, nil)

	input := dto.BeerBoxRequest{
		BeerID:   "1",
		Currency: "CLP",
		Quantity: "10",
	}
	beerBox, err := BoxBeerPriceById(repositoryMock, currencyMock, input)
	assert.Nil(t, err)
	assert.IsType(t, dto.BeerBoxResponse{}, beerBox)
	assert.Equal(t, int64(99809), beerBox.PriceTotal)
}

func TestBoxBeerPriceById_without_quantity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks_respository.NewMockBeerRepository(ctrl)
	repositoryMock.EXPECT().Get(gomock.Any()).Return(&domain.Beer{
		ID:       1,
		Name:     "Golden",
		Brewery:  "Kross",
		Country:  "Chile",
		Price:    10.5,
		Currency: "EUR",
	}, nil)
	currencyMock := mocks_respository.NewMockCurrencyRepository(ctrl)
	currencyMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(0.883155, 839.497922, nil)

	input := dto.BeerBoxRequest{
		BeerID:   "1",
		Currency: "CLP",
	}
	beerBox, err := BoxBeerPriceById(repositoryMock, currencyMock, input)
	assert.Nil(t, err)
	assert.IsType(t, dto.BeerBoxResponse{}, beerBox)
	assert.Equal(t, int64(59885), beerBox.PriceTotal)

}

func TestBoxBeerPriceById_with_equal_currency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks_respository.NewMockBeerRepository(ctrl)
	repositoryMock.EXPECT().Get(gomock.Any()).Return(&domain.Beer{
		ID:       1,
		Name:     "Golden",
		Brewery:  "Kross",
		Country:  "Chile",
		Price:    10.5,
		Currency: "EUR",
	}, nil)
	currencyMock := mocks_respository.NewMockCurrencyRepository(ctrl)

	input := dto.BeerBoxRequest{
		BeerID:   "1",
		Currency: "EUR",
		Quantity: "10",
	}
	beerBox, err := BoxBeerPriceById(repositoryMock, currencyMock, input)
	assert.Nil(t, err)
	assert.IsType(t, dto.BeerBoxResponse{}, beerBox)
	assert.Equal(t, int64(105), beerBox.PriceTotal)

}
