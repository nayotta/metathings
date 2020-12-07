#! /bin/bash

cd ${GOPATH}/src/github.com/nayotta/metathings

for src in $(find proto -name "*.proto" -exec bash -c 'dirname {}' \; | sort | uniq); do
    dst=pbset/${src}

    if [ ! -f ${dst} ]; then
        mkdir -p ${dst}
    fi

    protoc \
        -I${src} \
        -I${GOPATH}/src/github.com/nayotta/metathings/vendor \
        -I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
        -I${GOPATH}/src \
        --include_imports \
        --descriptor_set_out=${dst}/service.pbset \
        $(ls ${src}/*.proto)
done
