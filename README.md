# library
镜像源代码仓库，使用了 Go 模板语言进行渲染，便于批量构建和管理

## 镜像

### 自制镜像

自制镜像的用法和标签，参考各个子目录

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

自制镜像列表可以访问如下地址获得:

```
https://acicn.guoyk.net/library/IMAGES.txt
```
    
### 外部镜像

除了自制镜像外，我还导入了常用的外部镜像，对应关系参考 `manifest.yml` 文件，以下为示例

```
k8s.gcr.io/ingress-nginx/controller    =>  acicn/ingress-nginx-controller
k8s.gcr.io/defaultbackend-amd64        =>  acicn/ingress-nginx-defaultbackend
jettech/kube-webhook-certgen           =>  acicn/ingress-nginx-kube-webhook-certgen
quay.io/external_storage/nfs-client-provisioner    =>    acicn/nfs-client-provisioner
```

外部镜像列表可以访问如下地址获得:

```
https://acicn.guoyk.net/library/MIRRORS.txt
```

### 镜像源

除了标准的 DockerHub 源之外，我在其他云厂商的公共镜像仓库上也创建了对等的源 

* 腾讯云: `ccr.ccs.tencentyun.com/acicn`
* 阿里云: `registry.cn-shenzhen.aliyuncs.com/acicn`

## Helm 仓库

鉴于官方的 stable 中央仓库已经停用，我复制了一些常用的 Helm 仓库到 `charts` 子目录下，并定期发布。

Helm 仓库所涉及到的镜像均会作为外部镜像一并导入 `acicn` 命名空间

地址如下:

```
https://acicn-guoyk-net.oss-accelerate.aliyuncs.com/charts
```

## 许可证

Guo Y.K., MIT License
