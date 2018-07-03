GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/servo/dummy
BIN=$(OUTPATH)/ai.metathings.driver.servo.dummy.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
