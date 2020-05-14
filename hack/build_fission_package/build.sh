#! /bin/bash

set -ex

TMPDIR=${TMPDIR:-"/tmp"}
METATHINGS_HOME="$(go env GOPATH)/src/github.com/nayotta/metathings"
BUILD="${TMPDIR}/metathings_fission_package"

rm -rf ${BUILD}
cp -r ${METATHINGS_HOME} ${BUILD}
cd ${BUILD}
go mod tidy
go mod vendor

zip -qr evaluator_plugin.zip .
