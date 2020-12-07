FROM {{.repo_base}}/centos:8

RUN yum install -y @{{.centos_module}} && \
    yum install -y python3-devel && \
    yum clean all

ADD scripts /opt/bin
