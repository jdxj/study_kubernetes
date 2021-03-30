# 配置 configMap 卷

1. 指定配置文件夹

```shell
# config 为该项目的 config 文件夹, 主要关注 config.yaml 文件
$ $ kubectl create configmap study-k8s-config --from-file=config
configmap/study-k8s-config created
```

2. 查看一下
    - `|` 表示该值是多行

```shell
$ kubectl get configmap study-k8s-config -o yaml
...
apiVersion: v1
data:
  README.md: |-
    # 配置 configMap 卷

    1. 指定配置文件夹

    ```shell
    $ kubeclt
    ```
  config.go: "package config\n\nimport (\n\t\"io/ioutil\"\n\t\"os\"\n\n\t\"gopkg.in/yaml.v2\"\n)\n\nvar
    (\n\tConfigIns Config\n)\n\ntype Config struct {\n\tDB    DB    `yaml:\"db\"`\n\tRedis
    Redis `yaml:\"redis\"`\n}\n\ntype DB struct {\n\tUser string `yaml:\"user\"`\n\tPass
    string `yaml:\"pass\"`\n\tHost string `yaml:\"host\"`\n\tPort string `yaml:\"port\"`\n}\n\ntype
    Redis struct {\n\tHost  string `yaml:\"host\"`\n\tPort  string `yaml:\"port\"`\n\tPass
    \ string `yaml:\"pass\"`\n\tDBNum int    `yaml:\"dbNum\"`\n}\n\nfunc Init(path
    string) error {\n\tfile, err := os.Open(path)\n\tif err != nil {\n\t\treturn err\n\t}\n\tdefer
    file.Close()\n\n\tdata, err := ioutil.ReadAll(file)\n\tif err != nil {\n\t\treturn
    err\n\t}\n\n\treturn yaml.Unmarshal(data, &ConfigIns)\n}\n"
  config.yaml: |-
    db:
      user: ""
      pass: ""
      host: ""
      port: ""
    redis:
      host: ""
      port: ""
      pass: ""
      dbNum: 0
  config_test.go: "package config\n\nimport (\n\t\"fmt\"\n\t\"testing\"\n\n\t\"gopkg.in/yaml.v2\"\n)\n\nfunc
    TestMarshalConfig(t *testing.T) {\n\tcfg := &ConfigIns\n\tdata, err := yaml.Marshal(cfg)\n\tif
    err != nil {\n\t\tt.Fatalf(\"err: %s\\n\", err)\n\t}\n\tfmt.Printf(\"%s\\n\",
    data)\n}\n"
kind: ConfigMap
metadata:
  creationTimestamp: "2021-03-30T13:00:59Z"
  managedFields:
  - apiVersion: v1
    fieldsType: FieldsV1
    fieldsV1:
      f:data:
        .: {}
        f:README.md: {}
        f:config.go: {}
        f:config.yaml: {}
        f:config_test.go: {}
    manager: kubectl-create
    operation: Update
    time: "2021-03-30T13:00:59Z"
  name: study-k8s-config
  namespace: default
  resourceVersion: "160638"
  uid: d14521e3-bc73-44e2-96be-962717c64cf4
```