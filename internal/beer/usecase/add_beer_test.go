package usecase

import (
	"testing"

	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
	mocks_respository "github.com/alemelomeza/test-ms-beer/internal/beer/infra/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddBeers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks_respository.NewMockBeerRepository(ctrl)
	repositoryMock.EXPECT().Save(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(nil)

	input := dto.BeerItemRequest{
		Name:     "Golden",
		Brewery:  "Kross",
		Country:  "Chile",
		Price:    10.5,
		Currency: "EUR",
	}

	err := AddBeer(repositoryMock, input)
	assert.Nil(t, err)
}
