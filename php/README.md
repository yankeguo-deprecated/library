# php

`acicn/php` 镜像基于 `centos:8` 镜像

## 标签

* `acicn/php:7.2-centos-8`
* `acicn/php:7.3-centos-8`
* `acicn/php:7.4-centos-8`
* `acicn/php:7.4-ubuntu-20.04`
* `acicn/php:7.4-pagoda

鉴于我司的 PHP 程序员水平实在是有限，不得已单独创建 `7.4-bgy` 标签

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* PHP 扩展 和 REMI 源 (CentOS 版本)

  默认激活 Aliyun REMI 源

  本镜像使用 REMI 源安装 php-fpm 和 PHP 扩展，内置 PHP 扩展可以查阅 Dockerfile 文件

* PHP 扩展 (Ubuntu 版本)

  Ubuntu 版本使用另外一种方式，即从源代码编译 PHP 和 PHP-FPM 以及各种扩展，更加精准可控

* 使用 `merge-env-to-ini` 工具和环境变量修改 `PHP FPM` 配置文件

    详细参考 https://github.com/acicn/merge-env-to-ini

    * 环境变量前缀 `PHPCFG_PHP_INI_` 修改 `/etc/php.ini` 文件

        比如

        `PHPCFG_PHP_INI_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php.ini` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

    * 环境变量前缀 `PHPCFG_PHP_FPM_CONF_` 修改 `/etc/php-fpm.conf` 文件

        比如

        `PHPCFG_PHP_FPM_CONF_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php-fpm.conf` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

    * 环境变量前缀 `PHPCFG_PHP_FPM_WWW_CONF_` 修改 `/etc/php-fpm.d/www.conf` 文件

        比如

        `PHPCFG_PHP_FPM_WWW_CONF_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php-fpm.d/www.conf` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

* `nginx`

    `nginx` 进程完全使用 `acicn/nginx` 的配置模式，详情参考 https://github.com/acicn/library/tree/latest/nginx

    额外的修改

    - 使用文件 `/etc/nginx/default.conf.d/php.conf` 增加了 PHP 的支持（也就是 `location ~ \.php$ {` 区块）

    - 允许使用文件 `/etc/nginx/default.fastcgi.d/*.conf` 扩充上述区块的配置

    - 允许使用环境变量 `NGXCFG_DEFAULT_PHP_EXTRA_CONF` 扩充上述区块的配置

    - 默认启用 PHP 框架模式，即使用 `/var/www/public/index.php` 来统一处理所有路由
        - `NGXCFG_SNIPPETS_ENABLE_SPA=true`
        - `NGXCFG_SNIPPETS_SPA_INDEX=/index.php?$query_string`

## 默认配置

* PHP 项目地址 `/var/www/public`
