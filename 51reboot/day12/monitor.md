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



#### 通信协议

* 以'\n'结尾的json格式
* 没有ack