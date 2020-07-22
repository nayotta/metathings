RM=$(shell which rm)
CLEAN_PATHS=./bin ./lib
DOCKER_EXE=$(shell which docker)
CUR_PATH=$(shell pwd)

all: \
	protos

clean:
	$(RM) -rf $(CLEAN_PATHS)

protos_from_docker:
	$(DOCKER_EXE) run --rm -v $(CUR_PATH):/go/src/github.com/nayotta/metathings nayotta/metathings-development /usr/bin/make -C /go/src/github.com/nayotta/metathings/pkg/proto

protos:
	$(MAKE) -C pkg/proto all
