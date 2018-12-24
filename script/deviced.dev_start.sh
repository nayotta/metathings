#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/nayotta/metathings/cmd/metathingsd/main.go deviced -c /etc/metathings/deviced.yaml
