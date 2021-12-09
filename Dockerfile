FROM golang
WORKDIR /go/src/app
COPY . .
RUN go mod download && go build -o /go/bin/app cmd/api/main.go
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]