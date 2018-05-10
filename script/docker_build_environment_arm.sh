#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t bigdatagz/metathings-protobuf-arm -f dockerfile/Dockerfile.protobuf-arm .
docker build -t bigdatagz/metathings-env-arm -f dockerfile/Dockerfile.env-arm .
