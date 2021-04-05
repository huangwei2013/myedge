

## 定义
### 实体定义
server：服务端，提供
- grpc 服务，基于 google grpc 包实现；用于和 client 通讯
- http 服务，基于 goframe 实现；用于提供 web 管理功能
- 存储服务，基于(远端)ectd 实现

region：终端的管理单位，为多个 term 的集合
- 一个 term 同一时刻只属于一个 region
- 未指定的 term，属于 default region

term：实际终端
- 按自己的 mid 提交数据(历史 mid 存储于 server，连接时指定)


## 开发
### pb文件

* 编译

在 service/pb 目录中，执行如：
```
protoc -I ./ *.proto --go_out=plugins=grpc:.
```

## FAQ

### 环境安装
#### protoc-gen-go 安装

在 GOPATH/bin/protoc-gen-go 目录下，执行

```
go get github.com/golang/protobuf/protoc-gen-go
```
(若是在 goland 下，在其 terminal 中执行)

### 常见错误
#### undefined: grpc.SupportPackageIsVersion6 
参看 [资料](https://www.icode9.com/content-4-706221.html)