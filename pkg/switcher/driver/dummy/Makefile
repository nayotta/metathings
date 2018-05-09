GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/switcher/dummy
BIN=$(OUTPATH)/ai.metathings.driver.switcher.dummy.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
