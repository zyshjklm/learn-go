## 第13次课

上课时间：2017.0826

### 后继课程内容

rpc

web

db

二进制



### 作业点评

有人添加采集项的函数是阻塞的。这会导致需要在主函数中go该函数。最好是底层实现并发。

```go
sched.AddMetric(CpuMetric, time.Second)
// 这里只是关心一次添加操作。不应该在这里阻塞。
// 应该在内部使用go 协程。而不是让调用方添加协程。
```



#### 网络的错误处理涉及到：

用户态，内核态，对方服务。路由器，交换机。

常见问题：

* 网络中的路由将你的包丢弃或者拥塞时，客户端会有什么表现？
* 拨网线服务会有什么反应？
* 服务端将程序停掉，客户端会有什么表现？



网络连接中，tcp并不是物理连接，而是虚拟的。当拨掉网线时，需要看操作系统是否通知上层应用。如果没有数据发送时，通常不会直接感知到网络问题。以为连接依然存在。



* 三次握手
  * c: syn。net.Dial()发起连接
  * s: ack+syn。net.Listen()监听连接请求
  * c: ack
* 4次挥手。双方都使用conn.Close()请求关闭
  * 主动方为c or s。
    * c:fin
    * s:ack
  * 被动方
    * s: fin
  * c: ack



**处理假死：**

* 心跳包。没事也要经常联系
* 超时机制。超过时间就认为异常了




### 用户自定义监控项

agent中，默认集成的采集项，都有对应的函数。因此直接在调度中添加就可以了。

但对于自定义的监控项，需要调用额外的脚本，而默认的add函数的第一个参数是一个没有参数的函数。

```go
sched.AddMetric(MemMetric, time.Second*5)

// 自定义监控项
sched.AddMetric(NewUserMetric("./userDef.py"), time.Second*90)
```

原本是直接给一个函数的，现在需要根据一个脚本做参数，再返回一个函数，因此要使用闭包来返回这个函数。

实现闭包的思路：

* 先实现一个getUserMetrics()函数，以脚本名为参数，并返回同默认监控项一样的值。
* MemMetric()是没有参数的，getUserMetrics()是脚本名为参数
* 转换的关键就是实现一个闭包NewUserMetric()，返回getUserMetrics()




注意：

自定义监控user.go中，执行完命令后，如果没有`cmd.Wait()`则会产生僵死进程。

```shell
while true; do sleep 3; ps ax | grep '(Python)' | wc -l;done
     267
     268
     269

# 或者
watch 'ps ax | grep defunct | wc -l'
```



#### 写配置文件。

常用的配置文件：ini，yaml，toml。

```shell
# toml
go get -v github.com/BurntSushi/toml
```



### kafka



#### 相关概念

[参考](https://www.cnblogs.com/xjh713/p/7388262.html)：

* broker：对应集群的节点
* topic：对应的是一类消息。即一个逻辑上的应用数据
* partition：对一个topic的数据按hash进行多个分块上。
  * 从而能分到不同的broker以提搞读写速度
  * 每个partition对应到boker上的目录。
* segment：partition目录下按一定的偏移进行切分文件。包涵index和log
* offset：是个独立的消息或日志。每次顺序加一。最终写入segment文件。
* replica：patition的副本数。
  * 多个副本之间有leader/follower关系。类似hadoop的hdfs
  * 生产者写leader。

zookeeper用来记录消费者的读取偏移offset值。

特色：

* 顺序读写。并发高吞吐。
* 数据消费用不会直接删除，另有控制机制。
* 消费需要按part进行分组以提高性能。
* 客户端只记录offset。



#### trans的网络处理

* 建立监听
* 接收新连接
* 从连接读取数据
* 反序列化成common.Metric



反序列化的2种方法

```go
// 一种是：定义变量，传递指针
var metric common.Metric
json.Unmarshal([]byte(line), &metric)

// 一种是：new一个指针，传递变量
metric := new(common.Metric)
json.Unmarshal([]byte(line), metric)
```

实验代码参考：transfer-exercise1.go



#### go读写kafka

kafka库：**github.com/Shopify/sarama**

code: trans/transfer.go



### ES: elasticsearch

时间序列数据库。

* 按日期分隔
* 定期清理策略

方便扩容。

### saver

从kafka读数据，写入到ES。

```go

cluster "github.com/bsm/sarama-cluster"
elastic "gopkg.in/olivere/elastic.v5"

// 设置日间格式。16年01月02日03时04分05秒
time.Now().Format("20160102 030405")	// 	12小时制
time.Now().Format("20160102 150405")	// 	24小时制
```



待优化点：

* 写索引库的时间：
  * 应该需要根据metric中的时间来确定，根据metric的timestamp得到
  * 而不是本机的时间。saver简化了处理，直接根据本机的时间由`time.Now()`生成
* 需要定时定量写ES。现在是每来一条数据写一次，写盘压力大。



系统两个指标：

* 吞吐量。将多条数据汇总到一次进行写入。
* 时延。



监控的架构。此处没有深究

* agent如何升级和管理
* kafka/es的集群管理；
  * 旧数据的定期清理。
  * 热点数据的二次聚合。
* transfer的负载均衡。
  * agent要就近发送，
  * transfer自身的负载均衡

