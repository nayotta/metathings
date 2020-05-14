#! /bin/bash

set -ex

TMPDIR=${TMPDIR:-"/tmp"}
METATHINGS_HOME="$(go env GOPATH)/src/github.com/nayotta/metathings"
EVALUATOR_PLUGIN="${METATHINGS_HOME}/contrib/evaluator_plugin"
BUILD="${TMPDIR}/metathings_fission_package"

cd ${METATHINGS_HOME}
go mod tidy
go mod vendor

rm -rf ${BUILD}
cp -r ${EVALUATOR_PLUGIN} ${BUILD}
cd ${BUILD}
go mod tidy
go mod vendor
rm -rf vendor/github.com/nayotta/metathings
cp -r ${METATHINGS_HOME} vendor/github.com/nayotta/metathings

zip -qr evaluator_plugin.zip .
