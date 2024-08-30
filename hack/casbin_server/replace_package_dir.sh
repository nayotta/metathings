#! /bin/bash

sed -i 's|go_package="./;proto";|go_package="github.com/casbin/casbin-server/proto";|g' vendor/github.com/casbin/casbin-server/proto/casbin.proto
