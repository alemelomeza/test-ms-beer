package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
	"github.com/alemelomeza/test-ms-beer/internal/beer/dto"
	"github.com/alemelomeza/test-ms-beer/internal/beer/usecase"
	"github.com/gorilla/mux"
)

type handler struct {
	beerRepository     domain.BeerRepository
	currencyRepository domain.CurrencyRepository
}

func (h *handler) SearchBeers(w http.ResponseWriter, r *http.Request) {
	beers, err := usecase.SearchBeers(h.beerRepository)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(beers)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *handler) SearchBeerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["beerID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	beer, err := usecase.SearchBeerById(h.beerRepository, id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(beer)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *handler) AddBeers(w http.ResponseWriter, r *http.Request) {
	var beerItemRequest dto.BeerItemRequest
	err := json.NewDecoder(r.Body).Decode(&beerItemRequest)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = usecase.AddBeer(h.beerRepository, beerItemRequest)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) BoxBeerPriceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	beerBoxRequest := dto.BeerBoxRequest{
		BeerID:   vars["beerID"],
		Currency: r.URL.Query().Get("currency"),
		Quantity: r.URL.Query().Get("quantity"),
	}

	boxprice, err := usecase.BoxBeerPriceById(h.beerRepository, h.currencyRepository, beerBoxRequest)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(boxprice)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func NewHandler(beerRespository domain.BeerRepository, currencyRespository domain.CurrencyRepository) *handler {
	return &handler{
		beerRepository:     beerRespository,
		currencyRepository: currencyRespository,
	}
}
