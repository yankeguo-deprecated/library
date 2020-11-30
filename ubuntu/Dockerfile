FROM {{.repo_base}}/minit:{{.minit_version}} AS minit

FROM ubuntu:{{.ubuntu_version}}

ENV PATH "/opt/bin:${PATH}"
ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
        apt-get install -y language-pack-zh-hans && \
        apt-get upgrade -y && \
        apt-get install -y tzdata ca-certificates vim curl procps net-tools unzip netcat && \
        rm -rf /var/lib/apt/lists/* && \
        ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ADD scripts /opt/bin

COPY --from=minit /minit /opt/bin/minit

CMD ["/opt/bin/minit"]
