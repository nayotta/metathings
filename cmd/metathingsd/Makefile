GO=$(shell which go)
SRC=$(shell ls *.go)
BIN=../../bin/metathingsd

all: build

build:
	$(GO) build -o $(BIN) $(SRC)
