FROM {{.repo_base}}/alpine:3.12

RUN apk-mirror-use-aliyun && apk add --no-cache nginx && \
    rm -rf /etc/nginx/conf.d /etc/nginx/default.conf.d && \
    mkdir -p /etc/nginx/conf.d /etc/nginx/default.conf.d /etc/nginx/default.root.conf.d

ADD nginx.conf      /etc/nginx/nginx.conf
ADD snippets        /etc/nginx/snippets
ADD default.conf    /etc/nginx/conf.d/default.conf

ADD index.html      /var/www/public/index.html

ADD minit.d /etc/minit.d

WORKDIR /var/www/public

ENV NGXCFG_SNIPPETS_SPA_INDEX /index.html