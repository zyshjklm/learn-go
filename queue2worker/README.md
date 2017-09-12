## queue2worker README.md



### 文件说明：

* common/data.go。
  * 数据结构。**是需要修改的部分**，主要涉及:
  * 具体的struct
  * 对数据的处理方法
* common/worker.go。worker消费Post过来的数据，默认通道长度为1。
* common/dispatcher.go
  * Dispatcher结构及Job的调度工作。
  * 结构中包括了接收到的消息队列，以及worker的队列
  * 初始化所有worker
  * 每当消息队列中有数据，则取出来，分派给某个worker
* server/server.go
  * 有接收的包体的最大限制。默认是1M。
* client/client.go
  * 使用json格式传数据。
  * 是**需要修改的部分**。
  * 需要根据data.go中的数据结构，实现发送数据的方法。



### 使用方法

```shell
# client 

cd queue2worker/client

go run client.go
2017/09/12 23:13:18 {"version":"1.0","token":"TOKEN12345678","data":[{"metric":"cpu","value":81,"time":1505229198},{"metric":"mem","value":887,"time":1505229198}]}
2017/09/12 23:13:18 {"version":"1.0","token":"TOKEN12345678","data":[{"metric":"cpu","value":847,"time":1505229198},{"metric":"mem","value":59,"time":1505229198}]}

## server 

cd queue2worker/client

MAX_QUEUE=1024 MAX_WORKER=20 go run server.go
2017/09/12 23:12:22 queue:1024, worker:20
2017/09/12 23:12:22 [dispatch] queue:0; pool:20

2017/09/12 23:13:18 [dispatch] queue:1; pool:20
2017/09/12 23:13:18 [dispatch] queue:0; pool:20
2017/09/12 23:13:18 [process]: &{cpu 81 1505229198}
2017/09/12 23:13:18 [process]: &{mem 887 1505229198}
2017/09/12 23:13:18 [dispatch] queue:1; pool:20
2017/09/12 23:13:18 [process]: &{cpu 847 1505229198}
2017/09/12 23:13:18 [dispatch] queue:0; pool:20
2017/09/12 23:13:18 [process]: &{mem 59 1505229198}
```

