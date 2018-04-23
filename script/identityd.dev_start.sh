#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/bigdatagz/metathings/cmd/metathingsd/main.go identityd -c /etc/metathings/identityd.yaml
