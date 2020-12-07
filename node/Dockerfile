FROM {{.repo_base}}/centos:8

ENV PATH "/opt/node/bin:${PATH}"

RUN curl -sSL -o node.tar.gz "{{.node_download_url}}" && \
    mkdir -p /opt/node && \
    tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
    rm -f node.tar.gz && \
    npm config set unsafe-perm true
