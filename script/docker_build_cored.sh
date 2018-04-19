#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t bigdatagz/cored -f dockerfile/Dockerfile.cored .
