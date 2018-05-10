#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t bigdatagz/metathings-protobuf -f dockerfile/Dockerfile.protobuf .
docker build -t bigdatagz/metathings-development -f dockerfile/Dockerfile.dev .
