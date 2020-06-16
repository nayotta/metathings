#! /bin/bash

set -ex

METATHINGS_HOME=${METATHINGS_HOME:-"`pwd`"}
BUILD=${BUILD_PATH:-"build"}

cp -r ${METATHINGS_HOME} ${BUILD}
cd ${BUILD}
if [ "x${HACK_BRANCH}" != "x" ]; then
    sed -i "s/latest/${HACK_BRANCH}/g" go.mod
fi
go mod tidy
go mod vendor
zip -qr evaluator_plugin.zip .
