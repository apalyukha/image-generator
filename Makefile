.PHONY: build
build:
	rm -rf build && mkdir build && go build -o build/img_gen -v ./cmd

.PHONY: run
run:
	go run cmd/main.go

.DEFAUL_GOAL := build
