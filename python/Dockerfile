FROM {{.repo_base}}/ubuntu:{{.ubuntu_version}}

RUN apt-mirror-use-aliyun && apt-get update && \
    apt-get install -y python3 python3-dev python3-pip && \
    rm -rf /var/lib/apt/lists/*
