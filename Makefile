NAME := remote

default: build

build:
	@go build \
	-o out/${NAME} main.go