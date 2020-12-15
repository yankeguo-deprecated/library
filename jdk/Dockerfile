FROM {{.repo_base}}/{{.upstream}}

# JDK
{{if .jdk_version }}
RUN curl -sSLo openjdk.tar.gz '{{.jdk_download_url}}' && \
    mkdir -p /opt/jdk && \
    tar -C /opt/jdk --strip-components 1 -xf openjdk.tar.gz && \
    rm -f openjdk.tar.gz && \
    {{if .centos_version}}
    yum install -y fontconfig && \
    yum clean all && \
    {{end}}
    :

ENV JRE_VERSION {{.jdk_version}}
ENV JDK_VERSION {{.jdk_version}}
ENV JAVA_VERSION {{.jdk_version}}

ENV JRE_HOME  /opt/jdk
ENV JDK_HOME  /opt/jdk
ENV JAVA_HOME /opt/jdk

ENV PATH "${JAVA_HOME}/bin:${PATH}"

ADD scripts/java-wrapper /opt/bin/java-wrapper
{{end}}

# Arthas
{{if .arthas_version }}
RUN mkdir -p /opt/arthas && \
    curl -sSLo arthas.zip '{{.arthas_download_url}}' && \
    unzip -d /opt/arthas arthas.zip && \
    rm -rf arthas.zip && \
    {{if .centos_version}}
    yum install -y telnet && \
    yum clean all && \
    {{end}}
    :

ADD scripts/as.sh /opt/bin/as.sh

ENV ARTHAS_VERSION {{.arthas_version}}
{{end}}

# Maven
{{if .maven_version}}
RUN curl -sSLo maven.tar.gz '{{.maven_download_url}}' && \
    mkdir -p /opt/maven && \
    tar -C /opt/maven --strip-components 1 -xf maven.tar.gz && \
    rm -f maven.tar.gz

ENV PATH "/opt/maven/bin:${PATH}"
{{end}}