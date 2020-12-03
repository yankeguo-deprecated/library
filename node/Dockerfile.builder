FROM {{.repo_base}}/node:{{.upstream}}

RUN apt-get update && \
    apt-get install -y build-essential && \
    rm -rf /var/lib/apt/lists/* && \
    npm install -g cnpm && \
    npm cache clean -f && \
    npm config set unsafe-perm true && \
    cnpm config set unsafe-perm true
