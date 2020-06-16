#!/bin/bash

set -eux

srcDir=${GOPATH}/src/$(basename ${SRC_PKG})

trap "rm -rf ${srcDir}" EXIT

# http://ask.xmodulo.com/compare-two-version-numbers.html
version_ge() { test "$(echo "$@" | tr " " "\n" | sort -rV | head -n 1)" == "$1"; }

if [ -d ${SRC_PKG} ]
then
    echo "Building in directory ${srcDir}"
    ln -sf ${SRC_PKG} ${srcDir}
elif [ -f ${SRC_PKG} ]
then
    echo "Building file ${SRC_PKG} in ${srcDir}"
    mkdir -p ${srcDir}
    cp ${SRC_PKG} ${srcDir}/function.go
fi

cd ${srcDir}

LDFLAGS="-X main.FissionConfigName=evaluator-plugin-config/evaluator-plugin.yaml"

# use vendor mode if the vendor dir exists when go version is greater
# than 1.12 (the version that fission started to support go module).
if [ -d "vendor" ] && [ ! -z ${GOLANG_VERSION} ] && version_ge ${GOLANG_VERSION} "1.12"; then
  go build -mod=vendor -buildmode=plugin -ldflags="${LDFLAGS}" -o ${DEPLOY_PKG} ./contrib/evaluator_plugin/*.go
else
  go build -buildmode=plugin -ldflags="${LDFLAGS}" -o ${DEPLOY_PKG} ./contrib/evaluator_plugin/*.go
fi
