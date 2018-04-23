#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/bigdatagz/metathings/cmd/metathingsd/main.go cored -c /etc/metathings/cored.yaml
