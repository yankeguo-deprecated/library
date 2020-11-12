FROM {{.repo_base}}/jdk:{{.jdk_version}}

RUN curl -sSL -o tomcat.tar.gz "{{.tomcat_download_url}}" && \
    mkdir -p /usr/local/tomcat && \
    tar -C /usr/local/tomcat --strip-components 1 -xf tomcat.tar.gz && \
    rm -f tomcat.tar.gz && \
    apt-get update && \
    apt-get install -y libtcnative-1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/tomcat

ENV PATH "/usr/local/tomcat/bin:${PATH}"

ADD minit.d /etc/minit.d
ADD scripts /opt/bin