# Deployment

## 减慢升级, 观察滚动升级过程

```shell
$ kubectl patch deployment api-dep -p '{"spec":{"minReadySeconds": 20}}'
```

## 触发升级, 修改镜像

```shell
$ kubectl set image deployment api-dep api-c=jdxj/api:0.5.0
```