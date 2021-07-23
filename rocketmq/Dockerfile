FROM {{.repo_base}}/{{.upstream}}

{{if (ne .component "console")}}
RUN curl -sSLo rocketmq.zip "{{.rocketmq_download_url}}" && \
    unzip rocketmq.zip && \
    rm -f rocketmq.zip && \
    mv {{.rocketmq_download_junk_dir}} /opt/rocketmq
{{end}}

WORKDIR /opt/rocketmq

{{if (ne .component "console")}}
ENV ROCKETMQ_HOME /opt/rocketmq
{{end}}

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

{{if (eq .component "console")}}
RUN curl -sSL -o console.jar {{.rocketmq_console_download_url}}
{{end}}
