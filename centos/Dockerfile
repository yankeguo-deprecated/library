FROM {{.repo_base}}/minit:{{.minit_version}} AS minit

FROM centos:{{.upstream_version}}

ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN yum update -y && \
    yum install -y glibc-langpack-zh epel-release tzdata ca-certificates vim curl procps net-tools unzip nc && \
    yum clean all && \
    echo "LANG=\"$LANG\"" > /etc/locale.conf && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p /opt/bin
ENV PATH "/opt/bin:${PATH}"

COPY --from=minit /minit /opt/bin/minit
CMD ["/opt/bin/minit"]
