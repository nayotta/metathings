#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/nayotta/metathings/cmd/metathingsd/main.go sensord -c /etc/metathings/sensord.yaml
