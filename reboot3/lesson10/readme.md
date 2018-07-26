## 10本周课程：



### https证书过期。

具体功能实现，参考：

* github.com/jkak/opstools/domain/cmd/tls/tls.go

具体使用示例，参考：

* github.com/jkak/opstools/domain/cmd/https.go



### TCP相关

#### 1）处理超时



```go
// tcptest/cmd/client/client.go
common.SetTimeout(conn, 3)

// tcptest/cmd/server/server.go
// simulate timeout
time.Sleep(10 * time.Second)
```

run

```shell
# go run main.go server

# time go run main.go client
client called
send...
recv...
2018/07/14 11:18:29 read tcp 127.0.0.1:65512->127.0.0.1:2000: i/o timeout
exit status 1
go run main.go client  0.88s user 0.42s system 32% cpu 3.958 total
```

