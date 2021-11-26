CLEAN_PATHS=./bin ./lib

all: \
	protos_from_docker

clean:
	rm -rf $(CLEAN_PATHS)

protos_from_docker:
	docker run --rm --entrypoint /go/src/github.com/nayotta/metathings/hack/protos/gen_gopb.sh -v $(CURDIR):/go/src/github.com/nayotta/metathings jaegertracing/protobuf

pbsets_from_docker:
	docker run --rm --entrypoint /go/src/github.com/nayotta/metathings/hack/protos/gen_pbset.sh -v $(CURDIR):/go/src/github.com/nayotta/metathings jaegertracing/protobuf
