FROM {{.repo_base}}/ubuntu:{{.ubuntu_version}}

RUN apt-get update && \
    apt-get install -y python3 python3-dev python3-pip python3-venv python-is-python3 && \
    rm -rf /var/lib/apt/lists/*

ADD scripts /opt/bin