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





#### 2）内容长度字段及buffer

**server2**

先使用了一个回写的`handleFunc()`，读取一次，并将内容写回。

进行交互是使用的`func handleNewFunc(c net.Conn) `函数。

这里涉及到批量读取到buf的技巧。我了解到的有2种方式：

* 使用bufio.NewReader(conn)方式。可以指定每次读多少字节。
* handleNewFunc示范的方式。


**client端发送**

* 发送前先计算待发送信息的长度。用int32表示长度。
* 先发送信息长度，再发送信息内容。
  * 长度通过`int2byte`转换为[]byte
  * 整个信息由`Packet()`函数进行封装。

**server端接收**

* 先读取一定的buf。
* 通过`unPacket()`对buf进行解包。
  * 先用`bytes2Int()`从前4个字节中读取信息长度msgLen
  * 然后将buf随后msgLen长度的内容当做正文。
  * 将正文写到chan中
* 移动偏移量，进行一次解包循环
* 其实这个方法不太严谨，因为buf只有1024字节。
  * 如果包的长度可能超过，则会出现问题
  * 示例的每上信息段都很短，所以逻辑上没有问题。



server2, client2

发送30次，接受是按buf获取的，可能一次获取到多个。

```shell
# go run main.go server2

server called
listenning...


conn from 127.0.0.1:50251
{"ID":0, "Name":"user-0"}
{"ID":1, "Name":"user-1"}
{"ID":2, "Name":"user-2"}

{"ID":9, "Name":"user-9"}
{"ID":10, "Name":"user-10"}
{"ID":11, "Name":"user-11"}

{"ID":28, "Name":"user-28"}
{"ID":29, "Name":"user-29"}
2018/07/14 13:55:22 EOF

# go run main.go client2
client called
0: write 25 data
1: write 25 data
2: write 25 data

9: write 25 data
10: write 27 data
11: write 27 data

29: write 27 data
```



