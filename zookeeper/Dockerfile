FROM {{.repo_base}}/jdk:11

RUN curl -sSL -o zookeeper.tar.gz "{{.zookeeper_download_url}}" && \
    mkdir -p /opt/zookeeper && \
    tar -C /opt/zookeeper --strip-components 1 -xvf zookeeper.tar.gz && \
    rm -rf zookeeper.tar.gz

ENV PATH "/opt/zookeeper/bin:${PATH}"

WORKDIR /opt/zookeeper

ADD scripts /opt/bin
ADD conf    /opt/zookeeper/conf
ADD minit.d /etc/minit.d

ENV ZKCFG_tickTime 2000
ENV ZKCFG_initLimit 10
ENV ZKCFG_syncLimit 5
ENV ZKCFG_dataDir /data
ENV ZKCFG_clientPort 2181
ENV ZKCFG_maxClientCnxns 1024

ENV ZKAUTOCFG_MYID 1
# ENV ZKAUTOCFG_SERVERS zoo1:2888:3888,zoo2:2888:3888,zoo3:2888:3888
