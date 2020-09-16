FROM powerman/alpine-runit-volume:v0.3.2
SHELL ["/bin/ash","-e","-o","pipefail","-x","-c"]

LABEL org.opencontainers.image.source https://github.com/Djarvur/allcups-itrally-2020-task

ENV VOLUME_DIR=/home/app/var/data
ENV SYSLOG_DIR=$VOLUME_DIR/syslog
VOLUME $VOLUME_DIR

WORKDIR /home/app

COPY . .
RUN ln -nsf "$PWD"/init/* /etc/service/
