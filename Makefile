NAME := remote

default: build

build:
	@go build \
	-o out/${NAME} main.go

# Installation instructions for goreleaser: https://goreleaser.com/install/
cross_compile:
	@goreleaser build --snapshot --rm-dist

release:
	@goreleaser release --rm-dist
