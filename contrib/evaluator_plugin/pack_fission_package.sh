#! /bin/bash

set -eux

cp -r contrib/evaluator_plugin build
cd build
go mod tidy
go mod vendor
zip -r evaluator_plugin.zip .
