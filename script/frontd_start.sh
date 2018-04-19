#! /usr/bin/env sh

/usr/local/bin/envoy -c /etc/envoy.front.yaml --service-cluster frontd
