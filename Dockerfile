FROM powerman/alpine-runit-volume:v0.3.2
SHELL ["/bin/ash","-e","-o","pipefail","-x","-c"]

ENV VOLUME_DIR=/data \
    SYSLOG_DIR=/data/syslog
VOLUME /data

WORKDIR /app

COPY . .
RUN ln -nsf /app/init/* /etc/service/
