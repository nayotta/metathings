#! /bin/bash

cd ${GOPATH}/src/github.com/nayotta/metathings

for pbd in $(find proto -name "*.proto" -exec bash -c 'dirname {}' \; | sort | uniq); do
    protoc \
        -I${pbd} \
        -I${GOPATH}/src/github.com/nayotta/metathings/vendor \
        -I${GOPATH}/src \
        --go_out=plugins=grpc:. \
        $(ls ${pbd}/*.proto)
done
