# redis

`acicn/redis` 基于 `acicn/ubuntu:20.04` ，从源代码编译 `redis`

## 标签

* `acicn/redis:5`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* 透明大页

    可以使用环境变量 `MINIT_THP` 设置透明大页，**需要特权运行**

    比如

    `MINIT_THP=never`

* 内核参数

    可以使用环境变量 `MINIT_SYSCTL` 设置内核参数，**需要特权运行**

    比如

    `MINIT_SYSCTL=vm.overcommit_memory=1`

* 基于环境变量的 `redis.conf` 渲染

    所有以 `REDISCFG_` 开头的环境变量都会以以下规则转换为 `redis.conf` 中的配置

    - `_` 转换为 `-`
    - 重复的配置项，使用 `__` 添加后缀来避免环境变量冲突

    比如，以下环境变量

    ``` 
    REDISCFG_protected_mode=no
    REDISCFG_always_show_logo=no
    REDISCFG_save__1=900 1
    REDISCFG_save__2=300 10
    REDISCFG_save__3=60 10000
    ```

    会渲染为 `redis.conf` 配置项

    ``` 
    protected-mode no
    always-show-logo no
    save 900 1
    save 300 10
    save 60  10000
    ```

## 默认配置

* 工作目录 `/data`
* 数据目录 `./`，即 `/data`

详细默认配置，参见 `Dockerfile` 中的 `ENV` 配置
