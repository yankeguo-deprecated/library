FROM {{.repo_base}}/{{.upstream}}

RUN apk add --no-cache beanstalkd

WORKDIR /data

ADD minit.d /etc/minit.d

ENV BEANSTALKD_OPTS ""