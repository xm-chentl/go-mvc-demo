## 说明

### 命令

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    userservice/userservice.proto

```

### 问题

生成：The import path must contain at least one forward slash ('/') character.
v1.3.5以下，支持go_package 不加 /;

