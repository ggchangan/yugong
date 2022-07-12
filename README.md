# yugong

## 快速开始
### 快速部署
#### 启动数据库
```shell
cd deployments/docker && docker-compose up -d
```
#### Build & Run
```shell
make && ./apiserver -c configs/apiserver.yaml
```
### 构建