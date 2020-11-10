FROM {{.repo_base}}/ubuntu:20.04

RUN curl -sSL -o node.tar.gz "{{.node_download_url}}" && \
    mkdir -p /opt/node && \
    tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
    rm -f node.tar.gz

ENV PATH "/opt/node/bin:${PATH}"