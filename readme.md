# Protobuf

Protobuf是一种用于序列化结构化数据的灵活，高效，自动化的机制。
能够将结构化数据序列化,可用于数据存储，通信协议等方面。
您甚至可以更新数据结构，而不会破坏已针对“旧”格式编译的已部署程序。

## 优势

Protobuf相对于JSON和XML具有以下优点：

- 简洁
- 体积小，小3到10倍
- 速度快，快20到100倍
- 数据结构清晰
- 跨语言：编译生成各门编程语言使用的数据访问类

## 安装编译器

下载地址 ： https://github.com/protocolbuffers/protobuf/releases

## Protobuf Golang编译器

源码地址 ： https://github.com/golang/protobuf

> go install github.com/golang/protobuf/protoc-gen-go

## 编译生成各门编程语言使用的数据访问类

```shell script

# 编译生成go语言使用的数据访问类
protoc -I. -I%GOPATH%/src --go_opt=paths=source_relative --go_out=. ./go/proto/*.proto

# 编译生成c++语言使用的数据访问类
protoc -I. -I%GOPATH%/src --cpp_out=. ./cpp/proto/*.proto

```
