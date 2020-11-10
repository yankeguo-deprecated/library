# node

`acicn/node` 基于 `acicn/ubuntu:20.04`

## 标签

* `acicn/node:10`
* `acicn/node:12`
* `acicn/node:14`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    
    详细参考 https://github.com/acicn/minit

## 默认配置

* 安装目录 `/opt/node`
* `/opt/node/bin` 已经加入 `$PATH` 环境变量