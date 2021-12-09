package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/alemelomeza/test-ms-beer/internal/beer/domain"
)

type mySQLRepository struct {
	db *sql.DB
}

func (mr *mySQLRepository) List() ([]domain.Beer, error) {
	beers := make([]domain.Beer, 0)
	res, err := mr.db.Query("SELECT * FROM beers")
	if err != nil {
		return beers, err
	}
	defer res.Close()

	for res.Next() {
		var beer domain.Beer
		err := res.Scan(&beer.ID, &beer.Name, &beer.Brewery, &beer.Country, &beer.Price, &beer.Currency)
		if err != nil {
			return beers, err
		}
		beers = append(beers, beer)
	}

	return beers, nil
}

func (mr *mySQLRepository) Get(ID int64) (*domain.Beer, error) {
	beer := new(domain.Beer)
	err := mr.db.QueryRow("SELECT * FROM beers WHERE id = ?", ID).Scan(&beer.ID, &beer.Name, &beer.Brewery, &beer.Country, &beer.Price, &beer.Currency)
	if err != nil {
		return beer, err
	}
	return beer, nil
}

func (mr *mySQLRepository) Save(name string, brewery string, country string, price float64, currency string) error {
	sql := fmt.Sprintf("INSERT INTO `beers` (`brewery`, `country`, `currency`, `name`, `price`) values ('%s', '%s', '%s', '%s', %f)", brewery, country, currency, name, price)
	_, err := mr.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func NewMySQLRepository(user, password, dbhost, dbname string) *mySQLRepository {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+dbhost+":3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mySQLRepository{
		db: db,
	}
}
