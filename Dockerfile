FROM powerman/alpine-runit-volume:v0.3.2

ENV VOLUME_DIR=/data \
    SYSLOG_DIR=/data/syslog
VOLUME /data

WORKDIR /app

COPY . .
RUN set -x -e -o pipefail; \
    ln -nsf /app/init/* /etc/service/
