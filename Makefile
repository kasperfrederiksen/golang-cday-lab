run: 
	go run main.go

test: 
	go test ./...

build:
	go build -o bin/movies-api main.go