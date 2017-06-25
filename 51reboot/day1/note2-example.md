## go 语言特性

特点：

- 静态编译
- 垃圾回收
- 简洁的符号和语法
- 基于CSP的并发模型。
- 高效简单的工具链
- 丰富的标准库



### why golang

* 编译型，运行速度快
* 静态编译，没有依赖
* 天生支持并发，充分利用多核
* 大厂支持，有后台



### 服务器例子

```shell
vim server/server.go
go run server/server.go

## access
curl "http://localhost:8080/"
Hello, golang!
```



### c10k问题

```shell
vim c10k/c10k.go
go run c10k/c10k.go

while true;do curl "http://localhost:8080/";done
2017-06-18 20:15:52.464665102 +0800 CST
2017-06-18 20:15:52.475527952 +0800 CST
2017-06-18 20:15:52.487310341 +0800 CST
2017-06-18 20:15:52.496893867 +0800 CST
2017-06-18 20:15:52.50622846 +0800 CST
2017-06-18 20:15:52.520306789 +0800 CST
2017-06-18 20:15:52.529529313 +0800 CST
```



### 并行与并发

* 并发concurrent并不是并行parallel。
* node.js具有并发的能力，但不能充分利用多核。
* 写出一个能充分利用多核的程序需要很深的系统编程积淀
* 得益于优秀的设计，go可以轻松写出跑满所有cpu的程序



### 应用

* docker。容器化技术
* Kubernetes. Google Borg的开源实现
* Etcd