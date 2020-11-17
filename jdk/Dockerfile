FROM {{.repo_base}}/ubuntu:20.04

RUN curl -sSLo openjdk.tar.gz '{{.jdk_download_url}}' && \
    mkdir -p /opt/jdk && \
    tar -C /opt/jdk --strip-components 1 -xf openjdk.tar.gz && \
    rm -f openjdk.tar.gz && \
    apt-get update && \
    apt-get install -y telnet fontconfig && \
    rm -rf /var/lib/apt/lists/* && \
    curl -sSLo arthas.deb '{{.arthas_download_url}}' && \
    dpkg -i arthas.deb && \
    rm -rf arthas.deb

ENV JRE_HOME  /opt/jdk
ENV JDK_HOME  /opt/jdk
ENV JAVA_HOME /opt/jdk
ENV PATH "${JAVA_HOME}/bin:${PATH}"
# Ubuntu 系统比较邪门，默认 so 库不是安装在 /usr/lib 里面的
ENV LD_LIBRARY_PATH "/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu:/usr/local/lib/x86_64-linux-gnu"

ADD scripts /opt/bin
