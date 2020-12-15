# elasticsearch

## 标签

* `acicn/elasticsearch:6.3.2-jdk-11-centos-8`

## 功能

* 内置 `minit`

    * 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    * 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    * 内建 WebDAV 服务器，便于输出调试文件
    
    详细参考 https://github.com/acicn/minit

* 内核参数

    默认会使用 `MINIT_SYSCTL` 设置内核参数，**需要特权运行**
    
     * `vm.max_map_count=262144`

    可以将环境变量 `MINIT_SYSCTL` 设置为空字符串禁用此功能

* 资源限制

    默认会使用 `MINIT_RLIMIT_XXXX` 设置资源限制，**需要特权运行**

     *  `MINIT_RLIMIT_MEMLOCK=unlimited`

    可以将环境变量 `MINIT_RLIMIT_MEMLOCK` 设置为空字符串禁用此功能

* 使用 `elasticsearch-tune.jar` Java Agent 允许 Elasticsearch 运行在 root 用户下

* 安装 `analysis-ik`

* 基于环境变量的 `elasticsearch.yml` 配置文件渲染

    所有以 `ESCFG_` 开头的环境变量都会以以下规则转换为 `elasticsearch.yml` 中的配置

    * `__` 转换为 `.`

    比如，以下环境变量

    ```
    ESCFG_path__data=/data/data
    ESCFG_path__logs=/data/logs
    ```

    会渲染为 `elasticsearch.yml` 配置项

    ```
    path.data: /data/data
    path.logs: /data/logs
    ```

* 基于环境变量的 `jvm.options` 渲染

    * `ESJVM_XSS`
    * `ESJVM_XMS`
    * `ESJVM_XMX`
    
    **注意，也可以使用标准的 `ES_JAVA_OPTS` 环境变量**

## 默认配置

* 工作目录 `/opt/elasticsearch`
* 安装目录 `/opt/elasticsearch`
* 配置目录 `/opt/elasticsearch/config`
* 数据目录 `/data/data`
* 日志目录 `/data/logs`
* 发现模式 `single-node`

详细默认配置，参见 `Dockerfile` 中的 `ENV` 配置
