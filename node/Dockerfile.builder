FROM {{.repo_base}}/node:{{.upstream}}

RUN apt-get update && \
    apt-get install -y build-essential && \
    rm -rf /var/lib/apt/lists/* && \
    npm install -g cnpm && \
    npm cache clean -f

RUN npm config set registry https://r.npm.taobao.org
