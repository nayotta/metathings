#! /usr/bin/env bash

set -e

cd $(dirname $0)/..

./script/metathingsd_build.sh
docker build -t nayotta/metathingsd -f dockerfile/Dockerfile.metathingsd .
