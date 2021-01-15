# library
镜像源代码仓库，使用了 Go 模板语言进行渲染，便于批量构建和管理

## 用法

具体镜像用法和标签，参考各个子目录

* 操作系统
    * [Ubuntu](https://github.com/acicn/library/tree/latest/ubuntu)
    * [Alpine](https://github.com/acicn/library/tree/latest/alpine)
    * [CentOS](https://github.com/acicn/library/tree/latest/centos)
    * [Debian](https://github.com/acicn/library/tree/latest/debian)
* 运行环境
    * [JDK](https://github.com/acicn/library/tree/latest/jdk)
    * [PHP](https://github.com/acicn/library/tree/latest/php)
    * [Nginx](https://github.com/acicn/library/tree/latest/nginx)
    * [Node](https://github.com/acicn/library/tree/latest/node)
    * [Tomcat](https://github.com/acicn/library/tree/latest/tomcat)
* 基础服务
    * [Beanstalk](https://github.com/acicn/library/tree/latest/beanstalk)
    * [Elasticsearch](https://github.com/acicn/library/tree/latest/elasticsearch)
    * [MySQL](https://github.com/acicn/library/tree/latest/mysql)
    * [Mongo](https://github.com/acicn/library/tree/latest/mongo)
    * [PostgreSQL](https://github.com/acicn/library/tree/latest/postgres)
    * [Redis](https://github.com/acicn/library/tree/latest/redis)
    * [Zookeeper](https://github.com/acicn/library/tree/latest/zookeeper)
    
## 外部镜像

除了自制镜像外，我还导入了常用的外部镜像，映射关系如下

```
k8s.gcr.io/ingress-nginx/controller    =>  acicn/ingress-nginx-controller
k8s.gcr.io/defaultbackend-amd64        =>  acicn/ingress-nginx-defaultbackend
jettech/kube-webhook-certgen           =>  acicn/ingress-nginx-kube-webhook-certgen
quay.io/external_storage/nfs-client-provisioner    =>    acicn/nfs-client-provisioner
```

## 国内源

除了标准的 DockerHub 源之外，我还对等地创建了 腾讯云 和 阿里云 两个源，分别在 `registry.cn-shenzhen.aliyuncs.com/acicn` 和 `ccr.ccs.tencentyun.com/acicn` 命名空间下

## 许可证

Guo Y.K., MIT License
