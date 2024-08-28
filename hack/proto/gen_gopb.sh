#! /bin/bash

WORKDIR=`pwd`
VENDOR=${WORKDIR}/vendor

for pbd in $(find assets/proto -name "*.proto" -exec bash -c 'dirname {}' \; | sort | uniq); do
    outdir=$(echo ${pbd}|sed 's|^[^/]*/||')
    mkdir -p ${outdir}
    echo ">>>>>>>>>>>>>>>>>>>>"
    echo ${pbd}
    protoc \
            -I${WORKDIR} \
            -I${VENDOR} \
            -I${pbd} \
            --go_out=. \
            --go_opt=module=github.com/nayotta/metathings \
            --go-grpc_out=. \
            --go-grpc_opt=module=github.com/nayotta/metathings \
            $(ls ${pbd}/*.proto)
    echo "<<<<<<<<<<<<<<<<<<<<<"
done
