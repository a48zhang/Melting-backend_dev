# 2023-Melting-backend
Backend Repository for Melting

By @[hazardous waste](https://github.com/a48zhang)

By @[Cg1028](https://github.com/Cg1028)

---

# Release Note 
* V1.6 HTTPS支持

# 部署
## 使用Docker
### 拉取镜像
```shell
docker pull a48zhang/melting:dev
```
### 运行

```shell
docker rm $(sudo docker stop $(sudo docker ps -f name=melting_backend -q)) # 可选 关闭现有容器
docker run -dp <PORT>:65000 --env-file <environment-variable-filename> --name=melting_backend a48zhang/melting:dev
```

### 停止容器

```shell
docker stop $(sudo docker ps -f name=melting_backend -q)
```

---

# 维护
## 环境变量

请将下述环境变量导入Docker容器，否则服务不会运行。

如无特别说明均为必填。

| names       | descriptions   |
|-------------|----------------|
| access_key  | 七牛云_access_key |
| secret_key  | 七牛云_secret_key |
| bucket_name | 七牛云_存储桶名称      |
| domain_name | 七牛云_域名         |
| MELT_ADDR   | 后端服务_域名和端口     |
| MELT_CERT   | 后端服务_TLS证书     |
| MELT_KEY    | 后端服务_TLS密钥     |
| user        | MySQL数据库_用户名   |
| pwd         | MySQL数据库_密码    |
| addr        | MySQL数据库_地址    |
| db          | MySQL数据库_数据库名称 |
| MONGO_URL   | MongoDB_连接URI  |

# TODO List
* 使用MongoDB存储resource/
* 优化图片上传
* 日志记录


