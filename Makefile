
GOPATH:=$(go env GOPATH)

.PHONY: init
init:
	go mod download

.PHONY: build
build: init
	go build -o app ./cmd/api

.PHONY: test
test:
	go test -v -race -cover ./internal/beer/usecase

.PHONY: coverprofile
coverprofile:
	go test -coverprofile coverage.out ./internal/beer/usecase

.PHONY: cover
cover: coverprofile
	go tool cover -html=coverage.out

.PHONY: docker-up
docker-up:
	docker-compose up -d --build
 
.PHONY: docker-down
docker-down:
	docker-compose down