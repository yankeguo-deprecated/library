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

* PHP 模块

    - PPA 源 `ondrej/php`

      默认引入 PPA 源 `ondrej/php`，提供众多的 PHP 扩展模块

      详细列表参阅 https://launchpad.net/~ondrej/+archive/ubuntu/php

    - 默认模块

      `apcu`, `mysql`, `redis`, `mongodb`, `curl`, `mbstring`, `xml`, `zip`, `memcache`

    - 额外模块安装脚本 `php-extension-install`

      如果默认模块不符合要求，可以使用该命令安装任何 APT 源中已有的 PHP 模块，并且会自动清理临时文件，缩减镜像尺寸。内部调用 `apt-get install -y php-XXXX`

      示例

      ```dockerfile
      FROM acicn/php:7.2
      RUN php-extension-install apcu pgsql
      ADD . /var/www
      ```

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

* `nginx`

    `nginx` 进程完全使用 `acicn/nginx` 的配置模式，详情参考 https://github.com/acicn/library/tree/latest/nginx

    额外的修改

    - 使用文件 `/etc/nginx/default.conf.d/php.conf` 增加了 PHP 的支持（也就是 `location ~ \.php$ {` 区块）

    - 允许使用文件 `/etc/nginx/default.fastcgi.d/*.conf` 扩充上述区块的配置

    - 默认启用 PHP 框架模式，即使用 `/var/www/public/index.php` 来统一处理所有路由
        - `NGXCFG_SNIPPETS_ENABLE_SPA=true`
        - `NGXCFG_SNIPPETS_SPA_INDEX=/index.php?$query_string`

## 默认配置

* PHP-FPM 配置目录 `/etc/php/${PHP_VERSION}/fpm`
* PHP 项目地址 `/var/www/public`
