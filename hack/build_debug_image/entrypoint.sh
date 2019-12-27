#! /bin/sh

service=${SERVICE}
config=${CONFIG}

/usr/local/bin/dlv --listen=:44444 --headless=true exec /usr/local/bin/metathingsd -- ${service} -c ${config}
