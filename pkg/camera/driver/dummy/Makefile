GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/camera/dummy
BIN=$(OUTPATH)/ai.metathings.driver.camera.dummy.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
