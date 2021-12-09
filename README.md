# test-ms-beer

## Problem
Bender is a fan of beers, and he wants to keep track of all the beers he tastes and how to calculate the price he needs to buy a case of a specific type of beer. For this he needs a REST API with this information that he will later share with his friends.
 
## Description
It is requested to create a REST API based on the definition found in the file [openapi.yaml](api/openapi.yaml) in the `api` folder.

## Architecture
![Architecture](https://kroki.io/c4plantuml/svg/eNplkMtuwkAMRffzFW5WqVRY9QNoo3bBAhAR6ygkBkadB9hOgb_vzOSBRHe2r-_xY8FSk3TWqBftGtO1CMV7VXgneJP5OQpKtBiE8s6CFgYJWl0fqbZw8ASCLDPLsz0iDX1KbZDYu_xAGl37Btl3CrIQfUBfBFv_IBBeuuDnefaqem8eOVxxSkL_Z0wHbvIb4688QsSPiLRL8j5Y1ddN8qYjQtfcH8ixAjxRi7HUePeLJNq7iNmimW542mvHyH1HtUJ9PO09Pa3-f3AZQDzdHOcuy_UqUBZBCN_-A-YpjTM=)

## Requirements

* Golang ~1.16
* Docker ~20.10
* docker-compose ~1.29
* Make

## Use

Download source code:

```
git clone https://github.com/alemelomeza/test-ms-beer.git
```

Run services:

```
cd test-ms-beer && make docker-up
```
Stop services:

```
cd test-ms-beer && make docker-down
```

Services:

| Name | Port | Method | Endpoint | Params |
|-------|-------|-------|-------|-------|
| app | :8080 | GET | /beers | |
| app | :8080 | POST | /beers | name, brewery, country, price, currency |
| app | :8080 | GET | /beers/ | beerID|
| app | :8080 | GET | /beers/ | beerID, currency, quantity |
| db | :3306 | | |

Examples:

```
# Query - /beers
curl --location --request GET 'http://localhost:8080/beers'

# Response
[
    {
        "id": 1,
        "name": "Golden",
        "brewery": "Kross",
        "country": "Chile",
        "price": 10.5,
        "currency": "EUR"
    }
]
```

```
# Query - /beers/{beerID}
curl --location --request GET 'http://localhost:8080/beers/1'

# Response
{
    "id": 1,
    "name": "Golden",
    "brewery": "Kross",
    "country": "Chile",
    "price": 10.5,
    "currency": "EUR"
}
```

```
# Query - /beers/{beerID}/boxprice
curl --location --request GET 'http://localhost:8080/beers/1/boxprice?currency=CLP&quantity=10'

# Response
{
    "priceTotal": 99809
}
```

```
# Query - /beers
curl --location --request POST 'http://localhost:8080/beers' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "name": "Golden",
    "brewery": "Kross",
    "country": "Chile",
    "price": 10.5,
    "currency": "EUR"
}'
```

Tests:

Unit tests:

```
cd test-ms-beer && make test
```

End to End:

Use [test-ms-beer.postman_collection.json](test/test-ms-beer.postman_collection.json) in the `test` folder.


Coverage:

```
cd test-ms-beer && make cover
```

## Roadmap

* Impllement logs
* Implement a gracefull shutdown
* Implement JWT for authentication
* Enchance the docker image with multistage building
* Implement a layer of cache