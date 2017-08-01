## miniFtp.go

client从服务器获取文件：`GET XXXX.file.name`

```shell
#### server 
# go run miniFtp.go
2017/07/30 23:27:17 root: ./
2017/07/30 23:27:17 root: ./
2017/07/30 23:27:23 cmd:GET, name:/a.txts
2017/07/30 23:27:23 open .//a.txts: no such file or directory
2017/07/30 23:27:30 cmd:GET, name:/a.txt

#### client
# echo 'GET /a.txts' | nc localhost 8021
open .//a.txts: no such file or directory%                                                                                        

# echo 'GET /a.txt' | nc localhost 8021
hello golang!
```

如果使用前面实现的client

```shell
# go build -o ./client ../server/client_5file.go
# ./client
#GET /a.txt
2017/07/30 23:30:03 write size: 23
2017/07/30 23:30:03 return content:
hello golang!
```



#### STORE file

client向服务端发送文件内容。这里展示的是一个不太优雅的实现，主要是想说明遇到的问题及解决的思路。

服务端读取逻辑：

```go
// 参考storeFile函数。
buf := make([]byte, 4096)
for {
    rdNum, err := rd.Read(buf)
    if err != nil || err == io.EOF {
        break
    }
    wrNum, err := fd.Write(buf[:rdNum])
    if err != nil || wrNum != rdNum {
        break
    }
  	log.Print("read num:", rdNum)
}   
```

客户端写数据：

```go
conn.Write([]byte("STORE " + filepath.Base(fileName) + "\n"))	
num, err := io.Copy(conn, fd)
log.Print("write num:", num)
```

上述代码遇到的问题是，client端已经使用Copy发送完数据，但server端也正常读完了数据，却没有在EOF时退出。双方都在等待对方发送消息。

**处理办法有2个：**

* **提供文件长度：**
  * ****client在发送完文件名后，再另起一行发送文件的长度。
  * 之后再发文件内容。
  * server根据长度来判断是否读完内容
* **使用CloseWrite方法：**
  * 将conn断言成*net.TCPConn
  * 使用*net.TCPConn的CloseWrite方法。显式关闭发送，不再写数据
  * 从而再发送一次EOF。server端能收到EOF了。

第二个方法是老师提示的。很神奇。

观察第二个方法的运行效果：

server

```shell
# go run miniFtp.go
2017/08/01 08:41:25 root: ./
2017/08/01 08:41:25 root: ./
2017/08/01 08:42:05 cmd:STORE, name:ftp.txt
2017/08/01 08:42:05 read num:10
2017/08/01 08:42:05 cont:hello ftp

2017/08/01 08:42:05 read err:EOF

```

Client:

```shell
# cat ../ftp.txt
hello ftp

# go build -o client ./client_ftpStore.go

# ./client ../ftp.txt
2017/08/01 08:42:05 write num:10
2017/08/01 08:42:05 after close write
2017/08/01 08:42:05 return content:OK

# cat ftp.txt
hello ftp

# md5 ftp.txt ../ftp.txt
MD5 (ftp.txt) = 99d06d4c442a7b7e609f94029377b586
MD5 (../ftp.txt) = 99d06d4c442a7b7e609f94029377b586

```

