FROM {{.repo_base}}/{{.upstream}}

ENV APOLLO_VERSION {{.apollo_version}}

RUN mkdir -p /apollo/portal && cd /apollo/portal && \
    curl -sSL -o apollo-portal.zip "{{.apollo_download_url_portal}}" && \
    unzip apollo-portal.zip && rm -f apollo-portal.zip && \
    mkdir -p /apollo/configservice && cd /apollo/configservice && \
    curl -sSL -o apollo-configservice.zip "{{.apollo_download_url_configservice}}" && \
    unzip apollo-configservice.zip && rm -f apollo-configservice.zip && \
    mkdir -p /apollo/adminservice && cd /apollo/adminservice && \
    curl -sSL -o apollo-adminservice.zip "{{.apollo_download_url_adminservice}}" && \
    unzip apollo-adminservice.zip && rm -f apollo-adminservice.zip

ADD minit.d /etc/minit.d