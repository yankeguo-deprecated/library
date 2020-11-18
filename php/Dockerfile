FROM {{.repo_base}}/merge-env-to-ini AS utils-merge-env-to-ini

FROM {{.repo_base}}/nginx AS utils-nginx

FROM {{.repo_base}}/ubuntu:{{.ubuntu_version}}

ENV PHP_VERSION {{.php_version}}

RUN apt-mirror-use-aliyun && apt-get update && \
    apt-get install -y nginx php-fpm && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /etc/nginx && \
    mkdir -p /etc/nginx/default.fastcgi.d

COPY --from=utils-merge-env-to-ini /merge-env-to-ini /opt/bin/merge-env-to-ini

COPY --from=utils-nginx /etc/nginx /etc/nginx
COPY --from=utils-nginx /etc/minit.d/nginx.yml /etc/minit.d/nginx.yml

ADD minit.d /etc/minit.d
ADD index.php /var/www/public/index.php
ADD php.conf /etc/nginx/default.conf.d/php.conf

ENV PHPCFG_FPM_POOL_WWW_CONF_www__user  "user = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__group "group = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__listen_owner "listen.owner = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__listen_group "listen.group = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__listen "listen = /run/php/php-fpm.sock"
ENV PHPCFG_FPM_PHP_FPM_CONF_global__pid "pid = /run/php/php-fpm.pid"
ENV PHPCFG_FPM_PHP_FPM_CONF_global__error_log "error_log = /dev/stderr"

ENV NGXCFG_DEFAULT_EXTRA_INDEX index.php

ENV NGXCFG_SNIPPETS_ENABLE_SPA  true
ENV NGXCFG_SNIPPETS_SPA_INDEX   "/index.php?\$query_string"

