.PHONY: build
build:
	rm -rf build && mkdir build && go build -o build/2-image-generator-server -v ./cmd

.PHONY: run
run: 
	go run cmd/main.go

.DEFAULT_GOAL := build