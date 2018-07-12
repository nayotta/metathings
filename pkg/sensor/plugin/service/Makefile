GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/plugin/sensor
BIN=$(OUTPATH)/ai.metathings.service.sensor.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
