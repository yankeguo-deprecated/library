# nginx

`acicn/nginx` 镜像基于 `acicn/alpine:3.12`，使用默认的最新的 nginx 版本

## 标签

* `acicn/nginx:alpine-3.12`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    
    详细参考 https://github.com/acicn/minit

* 默认配置文件渲染

    `acicn/nginx` 的配置文件渲染经过详细设计，按照以下格式组成

    * `/etc/nginx/nginx.conf`
        * 默认设置
        * `http` 段
            * 默认设置
            * 引用环境变量 `NGXCFG_NGINX_EXTRA_CONF`
            * 引用文件 `/etc/nginx/conf.d/*.conf`
    * `/etc/nginx/conf.d/default.conf`
        * 默认设置
        * 引用环境变量 `NGXCFG_DEFAULT_EXTRA_CONF`
        * 引用文件 `/etc/nginx/default.conf.d/*.conf`
        * `location /`
            * 引用环境变量 `NGXCFG_DEFAULT_ROOT_EXTRA_CONF`
            * 引用文件 `/etc/nginx/default.root.conf.d/*.conf`

    确信，通过这种设计，无论是对 `nginx.conf`，默认服务器，和默认服务器的 `location /` 配置都可以通过环境变量和文件注入轻松处理

* 预置脚本

    预置脚本存放在 `/etc/nginx/snippets` 目录

    * `cors_params.conf`

        包含了开放 `CORS` 所需的配置，可以使用环境变量启用

        `NGXCFG_SNIPPETS_ENABLE_CORS_PARAMS=true`

    * `spa.conf`

        包含针对类似于 `Vue` 的单页面应用所需的配置，可以使用环境变量启用

        `NGXCFG_SNIPPETS_ENABLE_SPA=true`

        使用环境变量

        `NGXCFG_SNIPPETS_SPA_INDEX=/other/dir/index.html`

        来指定 `/index.html` 以外的默认文件

    * `mute_head_root.conf`

        包含了关闭腾讯云默认健康检查日志打印所需的配置，可以使用环境变量启动

        `NGXCFG_SNIPPETS_ENABLE_MUTE_HEAD_ROOT=true`

    * `healthz.conf`

        /helathz 健康检查接口，直接返回 200 OK，可以使用环境变量启动

        `NGXCFG_SNIPPETS_ENABLE_HEALTHZ=true`

## 默认配置

默认 nginx 根目录 `/var/www/public`
工作目录 `/var/www/public`
