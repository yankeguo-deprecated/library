FROM {{.repo_base}}/minit:{{.minit_version}} AS minit

FROM alpine:{{.alpine_version}}

ENV PATH "/opt/bin:${PATH}"
ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN apk upgrade --no-cache && \
    apk add --no-cache coreutils tzdata ca-certificates vim curl && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    echo $TZ > /etc/TZ

ADD profile.d /etc/profile.d
ADD banner.minit.txt /etc/banner.minit.txt
ADD scripts /opt/bin

COPY --from=minit /minit /opt/bin/minit

CMD ["/opt/bin/minit"]
