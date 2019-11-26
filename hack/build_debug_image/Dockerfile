FROM alpine

COPY metathingsd /usr/local/bin/metathingsd
COPY entrypoint.sh /entrypoint.sh

ARG SERVICE
ARG CONFIG

RUN chmod +x /usr/local/bin/metathingsd && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
