FROM {{.repo_base}}/{{.upstream}}

{{if .centos_version}}
RUN yum install -y @{{.centos_module}}/build && \
    yum clean all
{{end}}

ADD scripts /opt/bin

{{if .install_devel}}
{{if .centos_version}}
RUN yum install -y make automake gcc gcc-c++ kernel-devel && \
    yum clean all
{{end}}
{{end}}

ADD pip.conf /root/.pip/pip.conf
