#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t nayotta/metathings-protobuf -f dockerfile/Dockerfile.protobuf .
docker build -t nayotta/metathings-development -f dockerfile/Dockerfile.dev .
