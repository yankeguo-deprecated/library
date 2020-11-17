# python

`Python` 镜像基于 `acicn/ubuntu:20.04` 和 `acicn/ubuntu:18.04`

## 标签

* `acicn/python:3.8` (基于 `acicn/ubuntu:20.04`)
* `acicn/python:3.6` (基于 `acicn/ubuntu:18.04`)

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    
    详细参考 https://github.com/acicn/minit

* `pip-mirror-use-aliyun`

    一个辅助脚本，执行后会写入 `~/.pip/pip.conf` 文件，使用阿里云 pypi 镜像

* `venv-wrapper`

    一个辅助脚本，会先尝试 `source venv/bin/activate` 然后再执行后续动作

    常见用法如下

    ```dockerfile
    FROM acicn/python:3.8

    WORKDIR /work

    ADD requirements.txt requirements.txt

    RUN python -m venv venv && \
        source venv/bin/activate && \
        pip install -r requirements.txt

    ADD . .

    CMD ["/opt/bin/minit", "--", "venv-wrapper", "python", "main.py" ]
    ```