package usecase

import (
	"testing"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	mocks_respository "github.com/alemelomeza/test-ms-beer/internal/beer/infra/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestSearchBeers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mocks_respository.NewMockBeerRepository(ctrl)
	repositoryMock.EXPECT().List().Return([]domain.Beer{
		{
			ID:       1,
			Name:     "Golden",
			Brewery:  "Kross",
			Country:  "Chile",
			Price:    10.5,
			Currency: "EUR",
		},
	}, nil)

	beers, err := SearchBeers(repositoryMock)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(beers))
}
