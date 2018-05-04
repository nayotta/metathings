GO=$(shell which go)
SRC=$(shell ls *.go)
BIN=../../../../lib/plugin/ai.metathings.service.echo.so

all: build

build:
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
