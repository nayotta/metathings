GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/servo/i2c_xiaor
BIN=$(OUTPATH)/ai.metathings.driver.servo.i2c_xiaor.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
