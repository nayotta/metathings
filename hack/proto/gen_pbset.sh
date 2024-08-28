#! /bin/bash

WORKDIR=`pwd`
VENDOR=${WORKDIR}/vendor

for pbd in $(find assets/proto -name "*.proto" -exec bash -c 'dirname {}' \; | sort | uniq); do
    outdir=pbset/$(echo ${pbd}|sed 's|^[^/]*/||'|sed 's|^[^/]*/||')
    mkdir -p ${outdir}
    echo ">>>>>>>>>>>>>>>>>>>>"
    echo ${pbd}
    protoc \
            -I${WORKDIR} \
            -I${VENDOR} \
            -I${pbd} \
            --include_imports \
            --descriptor_set_out=${outdir}/service.pbset \
            $(ls ${pbd}/*.proto)
    echo "<<<<<<<<<<<<<<<<<<<<<"
done
