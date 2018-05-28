GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/motor/dummy
BIN=$(OUTPATH)/ai.metathings.driver.motor.dummy.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
