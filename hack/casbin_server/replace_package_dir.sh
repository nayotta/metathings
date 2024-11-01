#! /bin/bash

SED="sed"
if [ "x$(uname -s)" == "xDarwin" ]; then
    SED="gsed"
    which $SED > /dev/null
    if [ $? == 1 ]; then
        echo "require gsed for macOS"
        exit 1
    fi
fi

$SED -i 's|go_package = "./;proto";|go_package = "github.com/casbin/casbin-server/proto";|g' vendor/github.com/casbin/casbin-server/proto/casbin.proto
