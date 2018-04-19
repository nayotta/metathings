#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t bigdatagz/identityd -f dockerfile/Dockerfile.identityd .
