FROM acicn/merge-env-to-ini AS utils-merge-env-to-ini

FROM {{.repo_base}}/nginx AS utils-nginx

FROM {{.repo_base}}/{{.upstream}}

ENV PHP_VERSION {{.php_version}}

{{if .centos_version}}
RUN rpm -i https://mirrors.aliyun.com/remi/enterprise/remi-release-8.rpm && \
    rm -rf /etc/yum.repos.d/remi*
ADD remi-aliyun.repo /etc/yum.repos.d/remi-aliyun.repo

RUN yum module enable -y php:remi-{{.php_version}} && \
    yum install -y nginx aspell aspell-en aspell-sk php-cli php-fpm php-dbg \
    php-bcmath php-dba php-enchant php-gmp php-gd php-imap php-intl php-ldap php-mbstring php-mcrypt \
    php-mysqlnd php-odbc php-opcache php-pdo php-pear php-pecl-apcu php-pecl-gearman php-pecl-geoip \
    php-pecl-imagick php-pecl-igbinary php-pecl-memcache php-pecl-memcached php-pecl-mongodb php-pecl-redis \
    php-pecl-rdkafka php-pecl-swoole php-pecl-yaml php-pecl-zip php-pgsql \
    php-process php-pspell php-recode php-snmp php-soap php-sodium php-tidy php-xml php-xmlrpc && \
    yum clean all && \
    rm -rf /etc/nginx && \
    mkdir -p /etc/nginx/default.fastcgi.d

ENV PHPCFG_PATH_PHP_FPM_CONF        /etc/php-fpm.conf
ENV PHPCFG_PATH_PHP_FPM_WWW_CONF    /etc/php-fpm.d/www.conf
ENV PHPCFG_PATH_PHP_INI             /etc/php.ini
{{end}}

{{if .ubuntu_version}}
RUN apt-get update && \
    apt-get install -y nginx build-essential autoconf cmake cmake-extras openssl libssl-dev libsodium-dev \
                       libxml2-dev libsqlite3-dev libgd-dev libxpm-dev libicu-dev libbz2-dev libcurl4-openssl-dev \
                       libenchant-dev libgmp-dev libldap2-dev libaspell-dev libreadline-dev libsnmp-dev \
                       libxslt1-dev libgeoip-dev imagemagick libevent-dev liblz4-dev \
                       librdkafka-dev libpq-dev libyaml-dev libzip-dev libtidy-dev libmcrypt-dev libonig-dev \
                       libmysqlclient-dev libpcre3-dev libpspell-dev libldap2-dev libsasl2-dev && \
    rm -rf /var/lib/apt/lists

RUN mkdir -p /opt/src/php && \
    curl -sSL -o php.tar.gz '{{.php_download_url}}' && \
    tar -xf php.tar.gz -C /opt/src/php --strip-components 1 && \
    rm -f php.tar.gz && \
    cd /opt/src/php && \
    ./configure --prefix=/opt --enable-fpm --enable-gd --with-external-gd \
                --with-config-file-path=/etc --with-config-file-scan-dir=/etc/php.d \
                --with-jpeg --with-webp --with-xpm --with-freetype --enable-bcmath --with-zlib \
                --with-pcre-jit --enable-sockets --enable-soap --with-pdo-mysql --with-mysql-sock \
                --with-mysqli --enable-mysqlnd --with-openssl --with-zip --enable-mbstring \
                --enable-intl --enable-pcntl --with-bz2 --enable-calendar --enable-sysvmsg \
                --enable-sysvsem --enable-sysvshm --with-curl --enable-ftp --with-enchant \
                --enable-exif --with-gmp --with-gettext --with-sodium --enable-dba --with-pspell \
                --with-readline --with-snmp --with-xsl --enable-shmop --with-xmlrpc \
                --with-tidy --with-pgsql --with-pdo-pgsql --enable-phpdbg \
                --enable-phpdbg-webhelper --enable-phpdbg-readline --enable-sigchild \
                --with-pear --with-ldap --with-ldap-sasl && \
    make && make install

ENV PATH "/opt/sbin:$PATH"

ADD scripts /opt/bin

RUN mkdir -p /opt/src/libmemcached &&
    curl -sSL -o libmemcached.tar.gz 'https://mirrors.guoyk.net/launchpad.net/libmemcached/1.0/1.0.18/_download/libmemcached-1.0.18.tar.gz' && \
    tar -xf libmemcached.tar.gz -C /opt/src/libmemcached --strip-components 1 && \
    rm -f libmemcached.tar.gz && \
    cd /opt/src/libmemcached && \
    ./configure --enable-libmemcachedprotocol --enable-hsieh_hash --enable-memaslap --enable-deprecated && \
    make && make install

RUN acicn-php-ext-install php-mcrypt 'https://mirrors.guoyk.net/pecl.php.net/get/mcrypt-1.0.4.tgz' && \
    acicn-php-ext-install php-igbinary 'https://mirrors.guoyk.net/pecl.php.net/get/igbinary-3.2.1.tgz' && \
    acicn-php-ext-install php-memcached 'https://mirrors.guoyk.net/pecl.php.net/get/memcached-3.1.5.tgz' \
        --enable-memcached \
        --enable-memcached-session \
        --enable-memcached-igbinary \
        --enable-memcached-json \
        --enable-memcached-protocol && \
    acicn-php-ext-install php-memcache 'https://mirrors.guoyk.net/pecl.php.net/get/memcache-4.0.5.2.tgz' && \
    acicn-php-ext-install php-mongodb 'https://mirrors.guoyk.net/pecl.php.net/get/mongodb-1.5.5.tgz' \
        --enable-mongodb && \
    acicn-php-ext-install php-redis 'https://mirrors.guoyk.net/pecl.php.net/get/redis-4.2.0.tgz' && \
    acicn-php-ext-install php-rdkafka 'https://mirrors.guoyk.net/pecl.php.net/get/rdkafka-5.0.0.tgz' \
        --with-rdkafka && \
    acicn-php-ext-install php-swoole 'https://mirrors.guoyk.net/pecl.php.net/get/swoole-4.2.13.tgz' && \
    acicn-php-ext-install php-geoip 'https://mirrors.guoyk.net/pecl.php.net/get/geoip-1.1.1.tgz' \
        --with-geoip && \
    acicn-php-ext-install php-imagick 'https://mirrors.guoyk.net/pecl.php.net/get/imagick-3.4.4.tgz'

ADD {{.php_version}}/etc /etc

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

ENV PHPCFG_PHP_INI_Date__timezone 'date.timezone = "Asia/Shanghai"'
ENV PHPCFG_PHP_FPM_WWW_CONF_www__user  "user = root"
ENV PHPCFG_PHP_FPM_WWW_CONF_www__group "group = root"
ENV PHPCFG_PHP_FPM_WWW_CONF_www__listen "listen = 127.0.0.1:9000"
ENV PHPCFG_PHP_FPM_CONF_global__pid "pid = /run/php-fpm.pid"
ENV PHPCFG_PHP_FPM_CONF_global__error_log "error_log = /dev/stderr"

ENV NGXCFG_DEFAULT_EXTRA_INDEX index.php

ENV NGXCFG_SNIPPETS_ENABLE_SPA  true
ENV NGXCFG_SNIPPETS_SPA_INDEX   "/index.php?\$query_string"
