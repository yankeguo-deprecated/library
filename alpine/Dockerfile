FROM {{.repo_base}}/minit:{{.minit_version}} AS minit

FROM alpine:{{.upstream_version}}

ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN sed -i "s/dl-cdn.alpinelinux.org/{{.apk_mirror}}/g" /etc/apk/repositories && \
    apk upgrade --no-cache && \
    apk add --no-cache coreutils tzdata ca-certificates vim curl && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    echo $TZ > /etc/TZ

RUN mkdir -p /opt/bin
ENV PATH "/opt/bin:${PATH}"

COPY --from=minit /minit /opt/bin/minit
CMD ["/opt/bin/minit"]
