

## 作业讲解

重点是理解tcp协议。因为读写socket与读写文件即有相同之处，也有不同之处。

### tcp协议

电脑是如何传输数据的？TCP/IP是事实的标准。

TCP，首先需要寻址，连接上双方。连接的方式：

* 专线
* 分时系统。分时是按包传输而不是整个文件，也就是分组来进行传输的。

有了分组之后，如何区分目标机器，需要一个编号。IP层。——快递员。

世界很大，传出去的数据可能会丢失，如何确保信息可靠传输。 ——tcp

TCP层具体的功能：

* 对丢失的包信息重传，可靠的传输。
* 对包进行编号确保顺序，有序。
* 控制速度，拥塞控制，有自主权进行控制。
* 面向连接，是虚的连接，而不是物理连接。通过地址和编号来确定

因此，在使用tcp时，虽然底层已经将socket实现成文件模式，但和文件也有不同之处：

* 极端情况下，可能每次只传很小一部分内容。因此在接收时，需要考虑接收到的数量。


* tcp数据是流式的，要在tcp上面实现新的协议时，需要确保有结束符来标识一些结束位置。比如如何确定请求端发送的请求包头结束了。
* 在读取数据时需要循环读取。


原来的方式

```go
buf := make([]byte, 1024)
n, _ := conn.Read(buf)
```

上面的接口，只能读一次，将数据存入buf，然后就返回读到的长度和错误信息返回。此时读到的长度，可能并不能符合一个完整的逻辑交互过程。

新的方式，

```go
r := bufio.NewReader(conn)
line, err := r.ReadString('\n')
```

这样可以循环读取，直到遇到某个分隔符。这种方式适合于服务端读取client的请求头。

对于二进制协议。则需要不同根据其协议本身进行解析。

server

```shell
go run homework/ftp_server.go
2017/08/06 15:48:21 new connection.
2017/08/06 15:48:21 start to copy...
2017/08/06 15:48:21 after copy...
2017/08/06 15:48:21 start to close()...
```

Client

```shell
echo 'GET ./note.md' | nc localhost 8021
## 直接显示文件内容。

echo 'GET ./note.md' | nc localhost 8021 > new.md

md5 new.md note.md
MD5 (new.md) = 6fe63d39c95ca7b95605fa2c72a5795e
MD5 (note.md) = 6fe63d39c95ca7b95605fa2c72a5795e

echo 'STORE ./store.md' | nc localhost 8021 < note.md

md5 store.md note.md
MD5 (store.md) = 6fe63d39c95ca7b95605fa2c72a5795e
MD5 (note.md) = 6fe63d39c95ca7b95605fa2c72a5795e
```



### 读文件

几种方式：

* 按行读取：适应处理。
* 按块读取：用于大文件。
* 一次读取：用于小文件。
* io.Copy

对于按块读取：

```go
buf := make([]byte, 4096)
for {
  n, err:= f.Read(buf)
  if err == io.EOF {
    break
  }
  conn.Write(buf[:n])
}
```

上面的代码，其实只需要一句话：

```go
io.Copy(conn, f)
```

第一个参数是Writer接口，需要实现Write方法。实现了Write方式的包括，socket, 文件句柄，

第二个参数是Reader接口，需要实现一个Read方法。实现了Reader的包括文件句柄，socket, connection, bufio。

写文件时，io.Copy(f, r)



## new tips

groupcache

*groupcache* 是一个分布式缓存 go 语言库,支持多节点互备热数据,有良好的稳定性和较高的并发性

kcp



