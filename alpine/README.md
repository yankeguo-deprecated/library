# alpine

## 标签

* `acicn/alpine:3.12`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* 切换源到 `Aliyun`

* 切换时区到 `CST`

* 切换语言到 `zh_CN.UTF-8`
