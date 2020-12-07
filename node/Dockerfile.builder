FROM {{.repo_base}}/node:{{.upstream}}

RUN yum install -y make automake gcc gcc-c++ kernel-devel && \
    yum clean all && \
    npm install -g cnpm && \
    npm cache clean -f && \
    cnpm config set unsafe-perm true
