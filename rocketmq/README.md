# rocketmq

## 标签

* `acicn/rocketmq:4.6.1-broker`
* `acicn/rocketmq:4.6.1-namesrv`

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* 内置 `Alibaba Arthas`

    **注意，使用 `Arthas` 调试可能需要为容器提供内核权限**

    可以直接执行 `as.sh` 启动

* 使用环境变量修改 `broker.conf`

    任何以 `RMQBROKERCFG_` 开头的环境变量，均会渲染进入 `broker.conf` 文件

    例如:

      ```
      RMQBROKERCFG_namesrvAddr = rocketmq-namesrv.common:9876
      ```

    常见参数请访问 https://github.com/apache/rocketmq/blob/73bb4b402a5da02faf9feded8f4d79dd1ba88348/docs/en/Operations_Broker.md

* `java-wrapper`

    镜像内置脚本 `java-wrapper` 用以启动 `Broker` 和 `Namesrv`

    - 支持 `JAVA_OPTS` 环境变量

         `JAVA_OPTS` 和 **任何以 `JAVA_OPTS_` 开头的环境变量**，都会被扩展到 Tomcat 执行命令上

         建议的用法:

         `JAVA_OPTS_HEAP` 用于堆配置参数

         `JAVA_OPTS_HEAP=-Xms1g -Xmx4g`

         `JAVA_OPTS_GC` 用于内存回收配置参数

         `JAVA_OPTS_GC=-XX:+UseG1GC`

         当然你也可以一股脑把所有参数都放在 `JAVA_OPTS` 环境变量里

    - 兼容旧的 `JAVA_MEMORY_MAX`, `JAVA_MEMROY_MIN`, `JAVA_XMX` 和 `JAVA_XMS` 环境变量

## 默认配置

* 工作目录 `/opt/rocketmq`
* 安装目录 `/opt/rocketmq`
* 数据目录 `/data`
