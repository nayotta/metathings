GO=$(shell which go)
RM=$(shell which rm)
METATHINGS_SRC=$(shell ls cmd/metathings/*.go)
METATHINGSD_SRC=$(shell ls cmd/metathingsd/*.go)
CLEAN_PATHS=./bin

build:
	$(GO) build -o bin/metathings $(METATHINGS_SRC)
	$(GO) build -o bin/metathingsd $(METATHINGSD_SRC)

clean:
	$(RM) -rf $(CLEAN_PATHS)

all: build
