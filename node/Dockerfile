FROM {{.repo_base}}/{{.upstream}}

{{if .node_version}}
RUN curl -sSL -o node.tar.gz "{{.node_download_url}}" && \
    mkdir -p /opt/node && \
    tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
    rm -f node.tar.gz

ENV PATH "/opt/node/bin:${PATH}"
ENV NODE_VERSION {{.node_version}}

RUN npm config set unsafe-perm true && \
    npm config set registry https://registry.npm.taobao.org
{{end}}

{{if .install_devel}}
{{if .centos_version}}
RUN yum install -y make automake gcc gcc-c++ kernel-devel git && \
    yum clean all
{{end}}
{{end}}

{{if .install_cnpm}}
RUN npm install -g cnpm && \
    npm cache clean -f && \
    cnpm config set unsafe-perm true
{{end}}
