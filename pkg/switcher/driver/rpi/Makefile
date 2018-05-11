GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/switcher/rpi
BIN=$(OUTPATH)/ai.metathings.driver.switcher.rpi.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
