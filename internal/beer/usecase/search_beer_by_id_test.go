package usecase

import (
	"testing"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
	mocks_respository "github.com/alemelomeza/test-ms-beer/internal/beer/infra/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestSearchBeerByID(t *testing.T) {
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

	beer, err := SearchBeerById(repositoryMock, 1)
	assert.Nil(t, err)
	assert.IsType(t, dto.BeerItemResponse{}, beer)
}
