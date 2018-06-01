GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/motor/l298p
BIN=$(OUTPATH)/ai.metathings.driver.motor.l298p.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
