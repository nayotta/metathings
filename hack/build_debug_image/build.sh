#! /bin/bash

tag=${tag:-"nayotta/metathingsd:v0.0.0-debug"}

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l" -o hack/build_debug_image/metathingsd cmd/metathingsd/main.go
GO111MODULE=off GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hack/build_debug_image/dlv $GOPATH/src/github.com/go-delve/delve/cmd/dlv
docker build --network host -t "${tag}" -f hack/build_debug_image/Dockerfile hack/build_debug_image
