FROM {{.repo_base}}/{{.upstream}}

RUN curl -sSL -o activemq.tar.gz "{{.activemq_download_url}}" && \
    mkdir -p /opt/activemq && \
    tar -C /opt/activemq --strip-components 1 -xf activemq.tar.gz && \
    mv -f /opt/activemq/conf /opt/activemq/conf.orig && \
    mkdir -p /opt/activemq/conf && \
    rm -rf /opt/activemq/data && \
    mkdir -p /data && \
    ln -sf /data /opt/activemq/data && \
    rm -f activemq.tar.gz

WORKDIR /opt/activemq

ADD minit.d /etc/minit.d

ENV PATH "/opt/activemq/bin:$PATH"

ENV ACTIVEMQ_VERSION {{.activemq_version}}
