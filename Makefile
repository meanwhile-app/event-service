all: build

run:
	go run main.go

dev:
	air

build:
	CGO_ENABLED=0 go build main.go

