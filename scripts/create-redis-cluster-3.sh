#!/bin/bash

set -eu

NAMESPACE="${NAMESPACE}"
NAME="${NAME}"
PVC_CLASS="${PVC_CLASS}"
PVC_SIZE="${PVC_SIZE}"

for ID in "1" "2" "3"; do
	
echo "Workload ${NAME}-c${ID}"

cat <<-EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: ${NAME}-c${ID}
  namespace: ${NAMESPACE}
spec:
  selector:
    k8s-app: ${NAME}-c${ID}
  ports:
    - name: redis
      port: 6379
      protocol: TCP
      targetPort: 6379
    - name: redis-cluster
      port: 16379
      protocol: TCP
      targetPort: 16379
  type: ClusterIP
EOF

CLUSTER_IP=$(kubectl get -n ${NAMESPACE} service ${NAME}-c${ID} -o go-template --template {{.spec.clusterIP}})

echo "Cluster IP Created: ${CLUSTER_IP}"

cat <<-EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ${NAME}-c${ID}-data
  namespace: ${NAMESPACE}
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ${PVC_CLASS}
  resources:
    requests:
      storage: ${PVC_SIZE}
EOF

cat <<-EOF | kubectl apply -f -
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: ${NAME}-c${ID}
  namespace: ${NAMESPACE}
spec:
  replicas: 1
  serviceName: ${NAME}-c${ID}
  selector:
    matchLabels:
      k8s-app: ${NAME}-c${ID}
  template:
    metadata:
      labels:
        k8s-app: ${NAME}-c${ID}
    spec:
      volumes:
        - name: vol-data
          persistentVolumeClaim:
            claimName: ${NAME}-c${ID}-data
      containers:
        - name: ${NAME}-c${ID}
          image: 'acicn/redis:5'
          env:
            - name: REDISCFG_cluster_enabled
              value: "yes"
            - name: REDISCFG_cluster_config_file
              value: "nodes.conf"
            - name: REDISCFG_cluster_announce_ip
              value: "${CLUSTER_IP}"
          volumeMounts:
            - name: vol-data
              mountPath: /data
EOF

done
