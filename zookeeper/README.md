# zookeeper

`acicn/zookeeper` 基于 `acicn/jdk:11`

## 标签

* `acicn/zookeeper:3.6.2`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* 基于环境变量的 `zoo.cfg` 渲染

    - 以 `ZKCFG_` 开头的环境变量都会以以下规则转换为 `zoo.cfg` 文件中的配置
        - `__` 替换为 `.`

    比如，以下环境变量

    ``` 
    ZKCFG_tickTime=2000
    ZKCFG_initLimit=10
    ZKCFG_syncLimit=5
    ZKCFG_autopurge__snapRetainCount=5
    ```

    会渲染为

    ``` 
    tickTime=2000
    initLimit=10
    syncLimit=5
    autopurge.snapRetainCount=5
    ```

    - 可以使用 `ZKAUTOCFG_MYID` 和 `ZKAUTOCFG_SERVERS` 来自动生成 `server.X` 配置

    假设要设置一个 3 机集群，3 个容器的访问地址分别为 `zoo1` , `zoo2` 和 `zoo3`

    可以分别设置以下环境变量

    ``` 
    # zoo1
    # zoo.cfg 如下
    ZKAUTOCFG_MYID=1
    ZKAUTOCFG_SERVERS=zoo1:2888:3888,zoo2:2888:3888,zoo3:2888:3888

    # zoo2
    # zoo.cfg 如下
    ZKAUTOCFG_MYID=2
    ZKAUTOCFG_SERVERS=zoo1:2888:3888,zoo2:2888:3888,zoo3:2888:3888

    # zoo2
    # zoo.cfg 如下
    ZKAUTOCFG_MYID=3
    ZKAUTOCFG_SERVERS=zoo1:2888:3888,zoo2:2888:3888,zoo3:2888:3888
    ```

    三个容器会分别写入不同的 `myid` 文件，并分别生成如下配置文件

    ``` 
    # zoo1
    # myid 自动写入 1
    # zoo.cfg 如下
    server.1=0.0.0.0:2888:3888
    server.2=zoo2:2888:3888
    server.3=zoo3:2888:3888

    # zoo2
    # myid 自动写入 2
    # zoo.cfg 如下
    server.1=zoo1:2888:3888
    server.2=0.0.0.0:2888:3888
    server.3=zoo3:2888:3888

    # zoo3
    # myid 自动写入 3
    # zoo.cfg 如下
    server.1=zoo1:2888:3888
    server.2=zoo2:2888:3888
    server.3=0.0.0.0:2888:3888
    ```

* 基于环境变量的 JVM 配置

    - 使用环境变量 `ZKJVM_OPTS` 来追加 Java 命令
    - 使用环境变量 `ZKJVM_XMS` 和 `ZKJVM_XMX` 来指定 `-Xms`, `-Xmx` 参数

## 默认配置

* 工作目录 `/opt/zookeeper`
* 数据目录 `/data`
