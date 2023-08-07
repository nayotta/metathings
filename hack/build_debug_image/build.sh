#! /bin/bash

set -x

registry=${registry:-"debug.registry.local:5000"}
tag=${tag:-"${registry}/metathingsd:v0.0.0-debug"}
goos=${goos:-"linux"}
goarch=${goarch:-"amd64"}

GOOS=${goos} GOARCH=${goarch} CGO_ENABLED=0 go build -gcflags "all=-N -l" -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" -o hack/build_debug_image/metathingsd cmd/metathingsd/main.go

if [ ! -f hack/build_debug_image/dlv ]; then
    (cd hack/build_debug_image; \
     git clone https://github.com/go-delve/delve.git; \
     cd delve; \
     go mod tidy; \
     go mod vendor; \
     docker run --rm -it -w /opt/delve -v `pwd`:/opt/delve golang:1.20-alpine sh -c "GOOS=${goos} GOARCH=${goarch} CGO_ENABLED=0 go build -o dlv ./cmd/dlv"; \
     mv dlv ..; \
     cd ..; \
     rm -rf delve; \
     )
fi

docker build --network host -t "${tag}" -f hack/build_debug_image/Dockerfile hack/build_debug_image
