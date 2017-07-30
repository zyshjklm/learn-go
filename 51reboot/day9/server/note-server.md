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



#### ver 3.2 for and routine, and channel

在ver3.1基础上，将server中阻塞的部分则匿名协程独立成函数worker，再使用协程。

* server启动一个协程后，只管接收请求，并将请求写入channel。
* worker从channel中读取连接，并进行处理。

执行效果：

```shell
# go run 3routine2.go &

# nc localhost 8021 & nc localhost 8021
[1] 68495
2017-07-30 20:06:42.477293737 +0800 CST: hello golang
[1]  + 68495 done       nc localhost 8021
2017-07-30 20:06:43.48029226 +0800 CST: hello golang
```

由于server中只有一个协程。所有多个请求到达时，在协程内部还是串行处理的。从日志的时间上也能看到相关1秒。

#### ver 3.3 for and routine

在ver3.2基础上:

* 将server中启动一个协程，改为每收到一个请求时启动一个协程。

- 不再使用通道，而是直接传递连接。

执行效果：

```shell
# go run 3routine3.go &

# nc localhost 8021 & nc localhost 8021
[1] 69107
2017-07-30 20:16:07.421593995 +0800 CST: hello golang
2017-07-30 20:16:07.42162057 +0800 CST: hello golang
[1]  + 69107 done       nc localhost 8021
```

这个版本达到了并发的需求，不过每次生成一个新的协程，处理完一个请求又销毁了。代价有些大。

比较好的方式是在3routine2.go的基础上，server中启动多个协程作为一个池子。这样即使只使用一个channel，也能较好的达到并发。



#### ver 4 resp html

基于3routine2.go，将返回的内容由字符串改为一个html。

执行效果：

```shell
# go run 4html.go &

# nc localhost 8021
HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

<html>
<body>
<h1 style="color:red">hello golang</h1>
</body>
</html>
```

另可以在浏览器中访问127.0.0.1:8021观察网页结果。

