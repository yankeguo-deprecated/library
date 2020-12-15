FROM {{.repo_base}}/{{.upstream}}

{{if .alpine_version}}
RUN apk add --no-cache beanstalkd
{{end}}

WORKDIR /data

ADD minit.d /etc/minit.d

ENV BEANSTALKD_OPTS ""