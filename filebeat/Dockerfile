FROM {{.repo_base}}/{{.upstream}}

RUN curl -sSL -o filebeat.tar.gz "{{.filebeat_download_url}}" && \
    mkdir -p /usr/share/filebeat && \
    tar xf filebeat.tar.gz -C /usr/share/filebeat --strip-components 1 && \
    rm -f filebeat.tar.gz

WORKDIR /usr/share/filebeat

ADD minit.d /etc/minit.d
