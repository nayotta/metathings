#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/nayotta/metathings/cmd/metathingsd/main.go identityd2 -c /etc/metathings/identityd2.yaml
