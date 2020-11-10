FROM {{.repo_base}}/jdk:{{.jdk_version}}

RUN curl -sSL -o tomcat.tar.gz "{{.tomcat_download_url}}" && \
    mkdir -p /opt/tomcat && \
    tar -C /opt/tomcat --strip-components 1 -xf tomcat.tar.gz && \
    rm -f tomcat.tar.gz && \
    apt-get update && \
    apt-get install -y libtcnative-1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /opt/tomcat

ENV PATH "/opt/tomcat/bin:${PATH}"

ADD minit.d /etc/minit.d
ADD scripts /opt/bin