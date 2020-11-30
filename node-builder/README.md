# node-builder

`acicn/node-builder` 基于 `acicn/ubuntu-builder:20.04`

## 标签

* `acicn/node-builder:8`
* `acicn/node-builder:10`
* `acicn/node-builder:12`
* `acicn/node-builder:14`

## 工作目录

`/workspace`

## 构建参数

* `BUILDER_UID` 输出目录要使用的 UID，一般使用 `$(id -u)`
* `BUILDER_GID` 输出目录要使用的 GID，一般使用 `$(id -g)`
* `BUILDER_USE_CNPM` 是否要使用 `cnpm`，默认为 `false`
* `BUILDER_SCRIPT` 要执行的 `package.json` 脚本

## 缓存目录
