# Apollo

携程 Apollo 是垃圾，别用

## 标签

* `1.3.0`

## 环境变量

* `EUREKA_INSTANCE_IP_ADDDRESS`

## 配置

参考内容

```
# /apollo/portal/config/apollo-env.properties

local.meta=http://localhost:8080
dev.meta=http://fill-in-dev-meta-server:8080
fat.meta=http://fill-in-fat-meta-server:8080
uat.meta=http://fill-in-uat-meta-server:8080
lpt.meta=${lpt_meta}
pro.meta=http://fill-in-pro-meta-server:8080

# /apollo/portal/config/application-github.properties

spring.datasource.url = jdbc:mysql://xxxxxxxx:3306/xxxxxxxx?characterEncoding=utf8
spring.datasource.username = xxxxxxxx
spring.datasource.password = xxxxxxxx

# /apollo/portal/config/app.properties

appId=100003173
jdkVersion=1.8

# /apollo/configservice/config/application-github.properties

spring.datasource.url = jdbc:mysql://xxxxxxxx:3306/xxxxxxxx?characterEncoding=utf8
spring.datasource.username = xxxxxxxx
spring.datasource.password = xxxxxxxx

# /apollo/configservice/config/app.properties

appId=100003172

# /apollo/adminservice/config/application-github.properties

spring.datasource.url = jdbc:mysql://xxxxxxxx:3306/xxxxxxxx?characterEncoding=utf8
spring.datasource.username = xxxxxxxx
spring.datasource.password = xxxxxxxx

# /apollo/adminservice/config/app.properties

appId=100003171
```