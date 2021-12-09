package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alemelomeza/test-ms-beer/internal/beer/infra/repository"
	"github.com/alemelomeza/test-ms-beer/internal/rest"
	"github.com/gorilla/mux"
)

func main() {
	beerRepository := repository.NewMySQLRepository(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	currencyRepository := repository.NewCurrencylaryerRepository(os.Getenv("CURRENCYLAYER_ACCESS_KEY"))
	beerHandler := rest.NewHandler(beerRepository, currencyRepository)

	r := mux.NewRouter()

	r.HandleFunc("/beers", beerHandler.SearchBeers).Methods("GET")
	r.HandleFunc("/beers", beerHandler.AddBeers).Methods("POST")
	r.HandleFunc("/beers/{beerID}", beerHandler.SearchBeerById).Methods("GET")
	r.HandleFunc("/beers/{beerID}/boxprice", beerHandler.BoxBeerPriceById).Methods("GET")

	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           r,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
