GO=$(shell which go)
SRC=$(shell ls *.go)
BIN=../../bin/metathingsd
DOCKERBIN=../../bin/metathingsd-docker

all: build

build:
	$(GO) build -o $(BIN) $(SRC)

build_docker:
	$(GO) build -o $(DOCKERBIN) $(SRC)
