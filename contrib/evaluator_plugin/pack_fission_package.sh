#! /bin/bash

set -ex

TMPDIR=${TMPDIR:-"/tmp"}
METATHINGS_HOME=${METATHINGS_HOME:-"`pwd`"}
BUILD="${TMPDIR}/build"

rm -rf ${BUILD}

cp -r ${METATHINGS_HOME} ${BUILD}

cd ${BUILD}

go mod tidy
go mod vendor

zip -qr evaluator_plugin.zip .

mv evaluator_plugin.zip ${METATHINGS_HOME}
