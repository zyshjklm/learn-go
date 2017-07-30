## about server side 



#### ver 1 single req 

只能响应一次请求。

```shell
# go run server1_single.go &
[1] 60252
# nc localhost 8021
hello golang
[1]  + 60252 done       go run server1_single.go

# go run server1_single.go &
[1] 60363

# telnet localhost 8021
Trying ::1...
Connected to localhost.
Escape character is '^]'.
hello golang
Connection closed by foreign host.
[1]  + 60363 done       go run server1_single.go

```



#### ver 2 for 

使用for循环，一直监控端口并响应请求。

```shell
# go run 2for.go &

# nc localhost 8021
2017-07-30 17:13:14.029475469 +0800 CST: hello golang

# nc localhost 8021
2017-07-30 17:13:14.870952913 +0800 CST: hello golang

# nc localhost 8021 & nc localhost 8021
[1] 64985
2017-07-30 17:24:39.182592385 +0800 CST: hello golang
[1]  + 64985 done       nc localhost 8021
2017-07-30 17:24:40.184448905 +0800 CST: hello golang

# nc localhost 8021 & nc localhost 8021 &
[1] 65065
[2] 65066
➜  server git:(master) ✗ 2017-07-30 17:26:43.766197592 +0800 CST: hello golang
[1]  - 65065 done       nc localhost 8021
➜  server git:(master) ✗ 2017-07-30 17:26:44.771000995 +0800 CST: hello golang
[2]  + 65066 done       nc localhost 8021

```

改变：

* 使用for循环，不断监控端口
* 将监控和处理工作独立成server函数。
* 使用了1秒的延时来模拟网络耗时。观察请求响应的滞后性
* 连接2个请求并压后台，能观察到2个请求虽然是同时发出来的，但响应却是顺序的。



#### ver 3.1 for and routine

使用for循环，并使用routine来并发响应请求。

Ver 2版本，请求是串行的。处理原则：**哪里阻塞go哪里**。

代码变化：

```go
// origin code
time.Sleep(1000 * time.Millisecond)
conn.Write([]byte(time.Now().String() + ": hello golang\n"))
conn.Close()

// new code
go func() {
    time.Sleep(1000 * time.Millisecond)
    conn.Write([]byte(time.Now().String() + ": hello golang\n"))
    conn.Close()
}()
```

执行效果：

```shell
# go run 3routine1.go &

# nc localhost 8021 & nc localhost 8021
[1] 66075
2017-07-30 17:34:03.769712967 +0800 CST: hello golang
2017-07-30 17:34:03.76973386 +0800 CST: hello golang
[1]  + 66075 done       nc localhost 8021

```

2个请求在1秒后同时返回。从打印的日志中可以看到处理时的时间点，版本2是间隔1秒，而本次的时间即在同一秒内。

