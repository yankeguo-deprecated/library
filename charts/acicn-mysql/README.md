# acicn-mysql

## Values.yaml

**使用 HostPath 绑定主机**

```yaml
workload:
  image:
    tag: "5.7"
  schedule:
    nodeSelector:
      kubernetes.io/hostname: "tke-test-worker-3"
config:
  rootPassword: "qwertyqwerty"
  custom: |-
    [mysqld]
    federated
    max_allowed_packet=1024m
storage:
  type: "HostPath"
  hostPath:
    path: "/data/test-mysql-data"
```

**使用新创建的 PersistentVolume**

```yaml
workload:
  image:
    tag: "5.7"
config:
  rootPassword: "qwertyqwerty"
  custom: |-
    [mysqld]
    federated
    max_allowed_packet=1024m
storage:
  type: "PersistentVolume"
  persistentVolume:
    class: "cbs"
    capacity: "100Gi"
```

**已有的 PersistentVolumeClaim**

```yaml
workload:
  image:
    tag: "5.7"
config:
  rootPassword: "qwertyqwerty"
  custom: |-
    [mysqld]
    federated
    max_allowed_packet=1024m
storage:
  type: "PersistentVolumeClaim"
  persistentVolumeClaim:
    name: "some-data"
```

**启用 NodePort 和 LoadBalancer**

```yaml
workload:
  image:
    tag: "5.7"
config:
  rootPassword: "qwertyqwerty"
  custom: |-
    [mysqld]
    federated
    max_allowed_packet=1024m
service:
  nodePort:
    enabled: true
  loadBalancer:
    annotations:
      service.kubernetes.io/qcloud-loadbalancer-internal-subnetid: subnet-xxxxxxxx
    enabled: true
storage:
  type: "PersistentVolume"
  persistentVolume:
    class: "cbs"
    capacity: "100Gi"
```

## Credits

Guo Y.K., MIT License
