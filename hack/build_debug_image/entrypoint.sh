#! /bin/sh

service=${SERVICE}
config=${CONFIG}
blocking=${BLOCKING}

dlv_opts=""
if [ "x${blocking}" == "x" ]; then
    dlv_opts="${dlv_opts} --continue"
fi

/usr/local/bin/dlv ${dlv_opts} --listen=:44444 --headless --accept-multiclient exec /usr/local/bin/metathingsd -- ${service} -c ${config}
