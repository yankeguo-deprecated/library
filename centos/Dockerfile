FROM {{.repo_base}}/minit:{{.minit_version}} AS minit

FROM centos:{{.centos_version}}

ENV PATH "/opt/bin:${PATH}"
ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN yum update -y && \
        {{if .fix_lang}} sed -e '/override_install_langs/ s/^#*/#/' -i /etc/yum.conf && yum reinstall -y glibc-common && \ {{end}}
        yum install -y {{if .fix_lang}}{{else}}glibc-langpack-zh{{end}} epel-release tzdata ca-certificates vim curl procps net-tools unzip nc && \
        yum clean all && \
        echo "LANG=\"$LANG\"" > /etc/locale.conf && \
        ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ADD scripts /opt/bin

COPY --from=minit /minit /opt/bin/minit

CMD ["/opt/bin/minit"]
