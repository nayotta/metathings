RM=$(shell which rm)
CLEAN_PATHS=./bin ./lib
DOCKER_EXE=$(shell which docker)
CUR_PATH=$(shell pwd)

all: \
	protos

clean:
	$(RM) -rf $(CLEAN_PATHS)

protos_from_docker:
	$(DOCKER_EXE) run --rm --entrypoint /go/src/github.com/nayotta/metathings/hack/protos/gen_gopb.sh -v $(CUR_PATH):/go/src/github.com/nayotta/metathings jaegertracing/protobuf
