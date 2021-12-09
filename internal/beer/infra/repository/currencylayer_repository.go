package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type currencylayerRepository struct {
	accessKey string
}

func (cr currencylayerRepository) Get(currencyBeer string, currencyBeerBox string) (float64, float64, error) {
	req, err := http.NewRequest("GET", "http://apilayer.net/api/live?access_key="+cr.accessKey+"&currencies="+currencyBeer+","+currencyBeerBox, nil)
	if err != nil {
		return 0, 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, 0, err
	}
	quotes := result["quotes"].(map[string]interface{})
	if quotes == nil {
		return 0, 0, fmt.Errorf("quotes not found")
	}
	var cBeer, cBeerBox float64
	if c, ok := quotes["USD"+currencyBeer]; ok {
		cBeer = c.(float64)
	}
	if c, ok := quotes["USD"+currencyBeerBox]; ok {
		cBeerBox = c.(float64)
	}
	if cBeer == 0 || cBeerBox == 0 {
		return 0, 0, fmt.Errorf("quotes not found")
	}

	return cBeer, cBeerBox, nil
}

func NewCurrencylaryerRepository(accessKey string) *currencylayerRepository {
	return &currencylayerRepository{
		accessKey: accessKey,
	}
}
