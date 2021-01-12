FROM {{.repo_base}}/{{.upstream}}

# JDK
{{if .jdk_version }}

{{if .centos_version}}
{{if ne .centos_version 8}}
BAD
{{end}}
RUN curl -sSLo jdk.rpm '{{.jdk_el8_rpm_download_url}}' && \
    yum install -y ./jdk.rpm fontconfig && \
    yum clean all && \
    rm -f jdk.rpm
{{end}}

{{if .debian_version}}
RUN curl -sSLo jdk.deb '{{.jdk_deb_download_url}}' && \
    apt-get update && \
    apt-get install -y ./jdk.deb fontconfig && \
    rm -rf /var/lib/apt/lists/* && \
    rm -f jdk.deb

ENV LD_LIBRARY_PATH "/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu:/usr/local/lib/x86_64-linux-gnu"
{{end}}

ADD scripts/java-wrapper /opt/bin/java-wrapper

ENV JRE_VERSION  {{.jdk_version}}
ENV JDK_VERSION  {{.jdk_version}}
ENV JAVA_VERSION {{.jdk_version}}
{{end}}

# Arthas
{{if .arthas_version }}
RUN mkdir -p /opt/arthas && \
    curl -sSLo arthas.zip '{{.arthas_download_url}}' && \
    unzip -d /opt/arthas arthas.zip && \
    rm -rf arthas.zip && \
    {{if .debian_version}}
    apt-get update && \
    apt-get install -y telnet && \
    rm -rf /var/lib/apt/lists/* && \
    {{end}}
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