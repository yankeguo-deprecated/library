# prometheus

为什么这帮人非要在容器内用非 root 用户，在容器内直接用 root 不好么？是不是脑子有问题？

## 标签

* `acicn/prometheus:2.25`

## 功能

* 使用环境变量 `PROMETHEUS_OPTS` 添加额外的启动参数

## 默认配置

* 配置文件 `/opt/prometheus/prometheus.yml`
* 数据目录 `/data`