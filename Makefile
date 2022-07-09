run:
	go run -tags development main.go
.PHONY: run

build:
	go build -tags development -o bin/game main.go

release:
	go build -o bin/game_release main.go
