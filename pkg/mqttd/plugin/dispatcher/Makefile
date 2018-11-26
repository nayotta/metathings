GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/plugin/mqttd
BIN=$(OUTPATH)/ai.metathings.dispatcher.mqttd.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
