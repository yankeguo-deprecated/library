FROM acicn/merge-env-to-ini AS utils-merge-env-to-ini

FROM {{.repo_base}}/nginx AS utils-nginx

FROM {{.repo_base}}/{{.upstream}}

ENV PHP_VERSION {{.php_version}}

{{if .centos_version}}
RUN rpm -i https://mirrors.aliyun.com/remi/enterprise/remi-release-8.rpm && \
    rm -rf /etc/yum.repos.d/remi*
ADD remi-aliyun.repo /etc/yum.repos.d/remi-aliyun.repo

RUN yum module enable -y php:remi-{{.php_version}} && \
    yum install -y nginx php-cli php-fpm php-dbg \
    php-bcmath php-dba php-enchant php-gd php-imap php-intl php-ldap php-mbstring php-mcrypt \
    php-mysqlnd php-odbc php-opcache php-pdo php-pear php-pecl-apcu php-pecl-gearman php-pecl-geoip \
    php-pecl-igbinary php-pecl-memcache php-pecl-memcached php-pecl-mongodb php-pecl-redis php-pecl-xdebug php-pgsql \
    php-process php-pspell php-recode php-snmp php-soap php-tidy php-xml php-xmlrpc && \
    yum clean all && \
    rm -rf /etc/nginx && \
    mkdir -p /etc/nginx/default.fastcgi.d

ENV PHPCFG_PATH_PHP_FPM_CONF        /etc/php-fpm.conf
ENV PHPCFG_PATH_PHP_FPM_WWW_CONF    /etc/php-fpm.d/www.conf
ENV PHPCFG_PATH_PHP_INI             /etc/php.ini
{{end}}

COPY --from=utils-merge-env-to-ini /merge-env-to-ini /opt/bin/merge-env-to-ini

COPY --from=utils-nginx /etc/nginx /etc/nginx
COPY --from=utils-nginx /etc/minit.d/nginx.yml /etc/minit.d/nginx.yml

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
