FROM nayotta/metathingsd

COPY ./third_party/mosquitto/mosquitto-plugin.yaml /etc/metathings/mosquitto-plugin.yaml
COPY ./third_party/mosquitto/mosquitto-plugin_start.sh /usr/local/sbin/mosquitto-plugin_start.sh
CMD /usr/local/sbin/mosquitto-plugin_start.sh
