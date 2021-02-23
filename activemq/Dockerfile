FROM {{.repo_base}}/{{.upstream}}

RUN curl -sSL -o activemq.tar.gz "{{.activemq_download_url}}" &&
    mkdir -p /opt/activemq &&
    tar -C /opt/activemq --strip-components 1 -xf activemq.tar.gz &&
    rm -f activemq.tar.gz

WORKDIR /opt/activemq

ENV PATH "/opt/activemq/bin:$PATH"

ENV ACTIVEMQ_VERSION {{.activemq_version}}
