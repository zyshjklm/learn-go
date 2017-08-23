## 监控架构

从第十二次课开始，写个监控的原型。这里只记录一小部分尝试工作。主要结构独立目录:

* ./51reboot/mini-monitor/



架构图，参考架构班。

http://51reboot.com/course/arch/



psutil库doc

* http://godoc.org/github.com/shirou/gopsutil



```shell
# go get -v github.com/shirou/gopsutil
# go get -v golang.org/x/sys/unix

### set hosts
# sudo scutil --set HostName jungle85
```



### agent

* 采集器
* 调度器
* 网络模块
  * 连接处理，提高稳定性
  * 发送控制，提高吞吐量



#### 数据格式



```go
type Metric struct {
  Metric string
  Endpoint string
  Tag []string
  Value float64
  Timestamp int64
}
```

传输涉及两个层面的东西

* 传输协议。即传输层。transport。如http.
* 数据格式。wire format。json, protobuf, ...



#### 通信协议

* 以'\n'结尾的json格式
* 没有ack



### 准备

```shell
go get -v github.com/shirou/gopsutil
```

从github上，可以看到从页面链接到:

`godoc.org/github.com/shirou/gopsutil`



### 使用gopsutil进行物理资源采集



#### cpu采集

cpu.Percent()

采集原理。在一个点采集了下cpu的使用率，再间隔一个时间点，进行一次采集，将两次采集的值取差值，再除以间隔时间。

第二个参数：是否采集所有cpu的使用率。

##### 实验第一个参数：

```shell
# go build cpu-interval
# time ./cpu-interval
2017-08-23 11:44:49.025802232 +0800 CST
2017-08-23 11:44:54.028347412 +0800 CST
39.25
2017-08-23 11:44:54.028401335 +0800 CST
./cpu-interval  0.00s user 0.01s system 0% cpu 5.015 total

# go run cpu-interval.go
2017-08-23 11:51:22.597902869 +0800 CST
2017-08-23 11:51:27.601835974 +0800 CST
12.993503248375813
2017-08-23 11:51:27.60191499 +0800 CST

```

实验第二个参数：

```shell
# go run cpu-per.go
2017-08-23 11:56:41.46468289 +0800 CST
2017-08-23 11:56:46.469867286 +0800 CST
[5.3419870194707935]
2017-08-23 11:56:46.470057751 +0800 CST
2017-08-23 11:56:51.474070306 +0800 CST
[11.799999999999999 2.19560878243513 9 2.4]
2017-08-23 11:56:51.474118431 +0800 CST

```

