FROM {{.repo_base}}/{{.upstream}}

RUN curl -sSLo elasticsearch.tar.gz "{{.elasticsearch_download_url}}" && \
    mkdir -p /opt/elasticsearch && \
    tar -C /opt/elasticsearch --strip-components 1 -xf elasticsearch.tar.gz && \
    rm -f elasticsearch.tar.gz && \
    /opt/elasticsearch/bin/elasticsearch-plugin install -b "{{.analysis_ik_download_url}}" && \
    curl -sSLo /opt/elasticsearch-tune-1.0.jar "{{.elasticsearch_tune_download_url}}"

WORKDIR /opt/elasticsearch

ADD config config
ADD minit.d /etc/minit.d

ENV MINIT_SYSCTL vm.max_map_count=262144
ENV MINIT_RLIMIT_MEMLOCK unlimited

ENV ESJVM_XSS 1m
ENV ESJVM_XMS 1g
ENV ESJVM_XMX 1g

ENV ESCFG_path__data /data/data
ENV ESCFG_path__logs /data/logs
ENV ESCFG_network__host 0.0.0.0
ENV ESCFG_http__port 9200
ENV ESCFG_discovery__type single-node
