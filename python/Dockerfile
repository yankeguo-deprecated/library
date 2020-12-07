FROM {{.repo_base}}/centos:8

RUN yum install -y @{{.centos_module}}/build && \
    yum clean all

ADD scripts /opt/bin
