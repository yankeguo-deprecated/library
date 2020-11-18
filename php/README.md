# php

`acicn/php` 镜像基于 `acicn/ubuntu` 镜像

## 标签

* `acicn/php:7.2` 基于 `acicn/ubuntu:18.04`
* `acicn/php:7.4` 基于 `acicn/ubuntu:20.04`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    
    详细参考 https://github.com/acicn/minit

* 使用 `merge-env-to-ini` 工具和环境变量修改 `PHP FPM` 配置文件

    详细参考 https://github.com/acicn/merge-env-to-ini

    * 环境变量前缀 `PHPCFG_FPM_PHP_INI_` 修改 `/etc/php/${PHP_VERSION}/fpm/php.ini` 文件

        比如

        `PHPCFG_FPM_PHP_INI_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php/${PHP_VERSION}/fpm/php.ini` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

    * 环境变量前缀 `PHPCFG_FPM_PHP_FPM_CONF_` 修改 `/etc/php/${PHP_VERSION}/fpm/php-fpm.conf` 文件

        比如

        `PHPCFG_FPM_PHP_FPM_CONF_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php/${PHP_VERSION}/fpm/php-fpm.conf` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

    * 环境变量前缀 `PHPCFG_FPM_POOL_WWW_CONF_` 修改 `/etc/php/${PHP_VERSION}/fpm/pool.d/www.conf` 文件

        比如

        `PHPCFG_FPM_POOL_WWW_CONF_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php/${PHP_VERSION}/fpm/pool.d/www.ini` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突


## 默认配置

* PHP-FPM 配置目录 `/etc/php/${PHP_VERSION}/fpm`

   包括

    - `php.ini`
    - `php-fpm.conf`
    - `pool.d/www.conf`

    默认使用环境变量修改 `pool.d/www.conf` 中的用户和组为 `root`

* PHP 项目地址 `/var/www/public`
