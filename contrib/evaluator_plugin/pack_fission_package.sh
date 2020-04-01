#! /bin/bash

set -ex

cp -r contrib/evaluator_plugin build
cd build
if [ "x${HACK_BRANCH}" != "x" ]; then
    sed -i "s/latest/${HACK_BRANCH}/g" go.mod
fi
go mod tidy
go mod vendor
zip -qr evaluator_plugin.zip .
