#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

docker build -t bigdatagz/metathingsd -f dockerfile/Dockerfile.metathingsd .
