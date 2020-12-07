FROM {{.repo_base}}/centos:8

RUN curl -sSLo openjdk.tar.gz '{{.jdk_download_url}}' && \
    mkdir -p /opt/jdk && \
    tar -C /opt/jdk --strip-components 1 -xf openjdk.tar.gz && \
    rm -f openjdk.tar.gz && \
    yum install -y telnet fontconfig && \
    yum clean all && \
    mkdir -p /opt/arthas && \
    curl -sSLo arthas.zip '{{.arthas_download_url}}' && \
    unzip -d /opt/arthas arthas.zip && \
    rm -rf arthas.zip

ENV JRE_HOME  /opt/jdk
ENV JDK_HOME  /opt/jdk
ENV JAVA_HOME /opt/jdk
ENV PATH "${JAVA_HOME}/bin:${PATH}"

ADD scripts /opt/bin
