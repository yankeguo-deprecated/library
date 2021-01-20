# acicn-mysql

## Values.yaml

```yaml
workload:
  image:
    tag: "5.7"
  schedule:
    nodeSelector:
      kubernetes.io/hostname: "tke-test-worker-3"
config:
  rootPassword: "qewrtyqwerty"
  custom: |-
    [mysqld]
    federated
    max_allowed_packet=1024m
storage:
  type: "HostPath"
  hostPath:
    path: "/data/test-mysql-data"
```