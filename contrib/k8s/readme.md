# Metathings with K8S

```bash
kubectl create -n metathings-storage configmap metathings-deviced-scripts --from-file=./deviced-upsert-mongodb-ttl-indexes.js
```
