#! /usr/bin/env sh

set -e

/usr/local/bin/metathingsd plugin mosquitto -c /etc/metathings/mosquitto-plugin.yaml
