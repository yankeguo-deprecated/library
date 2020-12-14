FROM acicn/merge-env-to-ini AS utils-merge-env-to-ini

FROM {{.repo_base}}/nginx AS utils-nginx

FROM {{.repo_base}}/centos:8

ENV PHP_VERSION {{.php_version}}

RUN yum module enable -y php:{{.php_version}} && \
    yum install -y nginx php-* && \
    yum clean all && \
    rm -rf /etc/nginx && \
    mkdir -p /etc/nginx/default.fastcgi.d

COPY --from=utils-merge-env-to-ini /merge-env-to-ini /opt/bin/merge-env-to-ini

COPY --from=utils-nginx /etc/nginx /etc/nginx
COPY --from=utils-nginx /etc/minit.d/nginx.yml /etc/minit.d/nginx.yml

ADD scripts /opt/bin
ADD minit.d /etc/minit.d
ADD index.php /var/www/public/index.php
ADD php.conf /etc/nginx/default.conf.d/php.conf

ENV PHPCFG_PHP_FPM_WWW_CONF_www__user  "user = root"
ENV PHPCFG_PHP_FPM_WWW_CONF_www__group "group = root"
ENV PHPCFG_PHP_FPM_WWW_CONF_www__listen "listen = 127.0.0.1:9000"
ENV PHPCFG_PHP_FPM_CONF_global__pid "pid = /run/php-fpm.pid"
ENV PHPCFG_PHP_FPM_CONF_global__error_log "error_log = /dev/stderr"

ENV NGXCFG_DEFAULT_EXTRA_INDEX index.php

ENV NGXCFG_SNIPPETS_ENABLE_SPA  true
ENV NGXCFG_SNIPPETS_SPA_INDEX   "/index.php?\$query_string"
