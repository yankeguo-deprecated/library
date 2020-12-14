FROM {{.repo_base}}/alpine:{{.alpine_version}}

RUN apk add --no-cache beanstalkd

WORKDIR /data

ADD minit.d /etc/minit.d

ENV BEANSTALKD_OPTS ""