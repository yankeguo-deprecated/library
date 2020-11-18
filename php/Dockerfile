FROM {{.repo_base}}/merge-env-to-ini AS utils-merge-env-to-ini

FROM {{.repo_base}}/ubuntu:{{.ubuntu_version}}

ENV PHP_VERSION {{.php_version}}

RUN apt-mirror-use-aliyun && apt-get update && \
    apt-get install -y php-fpm && \
    rm -rf /var/lib/apt/lists/*

COPY --from=utils-merge-env-to-ini /merge-env-to-ini /opt/bin/merge-env-to-ini

ADD minit.d /etc/minit.d

ENV PHPCFG_FPM_POOL_WWW_CONF_www__user  "user = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__group "group = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__listen_owner "listen.owner = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__listen_group "listen.group = root"
ENV PHPCFG_FPM_POOL_WWW_CONF_www__listen "listen = /run/php/php-fpm.sock"
ENV PHPCFG_FPM_PHP_FPM_CONF_global__pid "pid = /run/php/php-fpm.pid"
ENV PHPCFG_FPM_PHP_FPM_CONF_global__error_log "error_log = /dev/stderr"