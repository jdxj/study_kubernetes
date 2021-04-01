# Deployment

## 减慢升级, 观察滚动升级过程

```shell
$ kubectl patch deployment echo-dep -p '{"spec":{"minReadySeconds": 20}}'
deployment.apps/api-dep patched
```

## 触发升级, 修改镜像


```shell
$ kubectl set image deployment echo-dep echo-c=jdxj/echo:0.1.0
```
