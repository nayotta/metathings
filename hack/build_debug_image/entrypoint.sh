#! /bin/sh

service=${SERVICE}
config=${CONFIG}

/usr/local/bin/metathingsd ${service} -c ${config}
