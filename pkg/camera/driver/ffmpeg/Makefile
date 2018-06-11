GO=$(shell which go)
SRC=$(shell ls *.go)
OUTPATH=../../../../lib/driver/camera/ffmpeg
BIN=$(OUTPATH)/ai.metathings.driver.camera.ffmpeg.so

all: build

build:
	mkdir -p $(OUTPATH)
	$(GO) build --buildmode=plugin -o $(BIN) $(SRC)
