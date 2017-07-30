## client



#### 版本1 buffer 4096

```shell
# go run client1_buf4096.go
119.75.213.61:80
192.168.1.108:56383

write size: 18

return content:
HTTP/1.1 302 Moved Temporarily
Date: Sun, 30 Jul 2017 02:20:43 GMT
Content-Type: text/html
Content-Length: 215
Connection: Keep-Alive
Location: http://www.baidu.com/search/error.html
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

<html>
<head><title>302 Found</title></head>
<body bgcolor="white">
<center><h1>302 Found</h1></center>
<hr><center>pr-nginx_1-0-347_BRANCH Branch
Time : Thu Jul 27 12:28:21 CST 2017</center>
</body>
</html>
```

使用说明：

* 通过`conn = net.Dial("tcp", addr)`来建立连接
* 通过`conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))`发送请求
* 通过`conn.Read(buf)`读取返回内容，关注判断返回的状态值
* `string(buf[:n])`代表读到的内容，而不是使用整个buf





#### 版本2 buffer 128

分批次读取返回内容。

```shell
# go run client2_buf128.go
119.75.213.61:80
192.168.1.108:56383

write size: 18

return content:
HTTP/1.1 302 Moved Temporarily
Date: Sun, 30 Jul 2017 02:20:43 GMT
## ....

# time go run client2_buf128.go;
### go run client2_buf128.go  0.23s user 0.12s system 0% cpu 1:00.33 total
```

使用说明：

* 使用较小的buf来观察程序的退出过程
* 需要使用`if err != nil || err == io.EOF {`来判断读取结束。
* 如果每次打印err变量的值。开始一直是nil。
* 读完后，会一直等待服务端退出，退出时会发送EOF。
* 读取baidu.com的eof需要60秒时间。



#### 版本3 按行读取

每次读取一行内容。

```shell
# go run client3_line.go
119.75.213.61:80
192.168.1.108:56383

write size: 18

return content:
HTTP/1.1 302 Moved Temporarily
Date: Sun, 30 Jul 2017 02:20:43 GMT
## ....
```

使用说明：

- 按行读取和使用较小的buf时有同时的问题，需要等待远端eof来控制程序的退出
- 程序60秒后会退出，并打印`2017/07/30 11:06:10 EOF`
- ​

#### 版本4 io.Copy

直接进行Copy操作。

```shell
# go run client4_copy.go
```

使用说明：

* 使用Copy，依然有等待远端eof来控制程序的退出的问题。
* 最开始使用4096的方式，因为没有使用for循环。读一次后就退出了。虽然没有等待问题，但需要返回内容的长度比在buf长度小。不然会导致没读完。

