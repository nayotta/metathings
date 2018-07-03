GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/plugin/servo
BIN=$(OUTPATH)/ai.metathings.dispatcher.servo.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
