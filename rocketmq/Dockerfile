FROM {{.repo_base}}/{{.upstream}}

RUN curl -sSLo rocketmq.zip "{{.rocketmq_download_url}}" && \
    unzip rocketmq.zip && \
    rm -f rocketmq.zip && \
    mv {{.rocketmq_download_junk_dir}} /opt/rocketmq

WORKDIR /opt/rocketmq

ENV ROCKETMQ_HOME /opt/rocketmq

ADD {{.component}}/minit.d /etc/minit.d

{{if (eq .component "broker")}}

RUN mkdir -p /data && ln -sf /data /root/store

ADD broker.conf /opt/rocketmq/conf/broker.conf

ENV RMQBROKERCFG_brokerClusterName  DefaultCluster
ENV RMQBROKERCFG_brokerName         broker-a
ENV RMQBROKERCFG_brokerId           0
ENV RMQBROKERCFG_deleteWhen         04
ENV RMQBROKERCFG_fileReservedTime   48
ENV RMQBROKERCFG_brokerRole         ASYNC_MASTER
ENV RMQBROKERCFG_flushDiskType      ASYNC_FLUSH

{{end}}
