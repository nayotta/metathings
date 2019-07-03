#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/nayotta/metathings/cmd/metathingsd/main.go device_cloud -c /etc/metathings/device_cloud.yaml
