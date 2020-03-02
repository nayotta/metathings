#! /bin/bash

GO111MODULE=on go mod vendor
cp -r contrib/evaluator_plugin build
mv vendor build
cd build
zip -r evaluator_plugin.zip .
