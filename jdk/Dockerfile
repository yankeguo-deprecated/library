FROM {{.repo_base}}/ubuntu:20.04

RUN curl -sSLo openjdk.tar.gz '{{.jdk_download_url}}' && \
    mkdir -p /opt/jdk && \
    tar -C /opt/jdk --strip-components 1 -xf openjdk.tar.gz && \
    rm -f openjdk.tar.gz && \
    apt-get update && \
    apt-get install -y telnet && \
    rm -rf /var/lib/apt/lists/* && \
    curl -sSLo arthas.deb '{{.arthas_download_url}}' && \
    dpkg -i arthas.deb && \
    rm -rf arthas.deb

ENV JRE_HOME  /opt/jdk
ENV JDK_HOME  /opt/jdk
ENV JAVA_HOME /opt/jdk
ENV PATH "${JAVA_HOME}/bin:${PATH}"

ADD scripts /opt/bin
