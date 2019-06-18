#! /usr/bin/env sh

set -e

go run $GOPATH/src/github.com/nayotta/metathings/cmd/metathingsd/main.go plugin mosquitto -c /etc/metathings/mosquitto-plugin.yaml
