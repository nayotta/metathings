#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t bigdatagz/metathings-development -f dockerfile/Dockerfile.dev .
