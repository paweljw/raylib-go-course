run:
	go run -tags development cmd/main.go
.PHONY: run

build:
	go build -tags development -o bin/game cmd/main.go

release:
	go build -o bin/game_release cmd/main.go
