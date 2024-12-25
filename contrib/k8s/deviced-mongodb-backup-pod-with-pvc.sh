#!/bin/bash

kubectl run deviced-mongodb-backup-pod \
  --image="registry.cn-hangzhou.aliyuncs.com/nayotta/bitnami-mongodb:8.0.4-debian-12-r0" \
  --overrides='{
    "apiVersion": "v1",
    "spec": {
      "volumes": [
        {
          "name": "backup-dir",
          "persistentVolumeClaim": {
            "claimName": "metathings-deviced-backup-20241224"
          }
        }
      ],
      "containers": [
        {
          "name": "instance",
          "image": "registry.cn-hangzhou.aliyuncs.com/nayotta/bitnami-mongodb:8.0.4-debian-12-r0",
          "command": ["/bin/bash", "-c", "sleep infinity"],
          "volumeMounts": [
            {
              "mountPath": "/mnt/backup",
              "name": "backup-dir"
            }
          ]
        }
      ]
    }
  }'
