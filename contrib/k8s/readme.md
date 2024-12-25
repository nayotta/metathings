# Metathings with K8S

```bash
kubectl create configmap \
  -n metathings-storage \
  metathings-deviced-scripts \
  --from-file=./deviced-upsert-mongodb-ttl-indexes.js \
  --from-file=./deviced-mongodb-backup-daily.sh \
  --dry-run=client -o yaml | kubectl apply -f -
```
