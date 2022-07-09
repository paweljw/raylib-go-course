run:
	go run main.go
.PHONY: run

build:
	go build -tags development -o bin/game main.go
