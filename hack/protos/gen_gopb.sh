#! /bin/bash

cd ${GOPATH}/src/github.com/nayotta/metathings

for pbd in $(find proto -name "*.proto" -exec bash -c 'dirname {}' \; | sort | uniq); do
    protoc \
        -I${pbd} \
        -I${GOPATH}/src/github.com/nayotta/metathings/vendor \
        -I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
        -I${GOPATH}/src \
        --go_out=plugins=grpc:${pbd} \
        --validate_out=lang=go:${pbd} \
        $(ls ${pbd}/*.proto)
done
