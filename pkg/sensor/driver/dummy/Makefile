GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/sensor/dummy
BIN=$(OUTPATH)/ai.metathings.driver.sensor.dummy.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
