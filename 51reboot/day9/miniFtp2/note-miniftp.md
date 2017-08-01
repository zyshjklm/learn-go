## miniFtp2.go



#### STORE file

两端都使用io.Copy，发送端能从文件句柄中读到EOF，故能正常结束。但接收端却没有正常识别。还在等待。使用closeWrite才能正常结束。

另一个问题是，接收端的接收方式：

```go
num, _ := io.Copy(conn, fd)
// 修改为
num, _ := io.Copy(conn, conn)
```

两种方式的结果看上去是一样的。

```shell
# md5 ftp.txt ../ftp.txt
MD5 (ftp.txt) = 99d06d4c442a7b7e609f94029377b586
MD5 (../ftp.txt) = 99d06d4c442a7b7e609f94029377b586
```

**问题说明：**

* **问题1：两端都使用io.Copy的问题**
  * 发送端读的是文件，读到结束时会收到EOF，故client端正常结束
  * 但io.Copy不会主动调用close。故发送端不会发送EOF
  * 因此，接收端从发送端的Copy中是读不到EOF的。需要追加写个EOF
* **问题2：接收端读bufio.Reader和net.Conn的一致性**
  * 从server/readBuf.go可以看到，读文件时，使用Reader和fd是有差别的。
    * 无论什么时候读文件，都是有数据的。至少能得到EOF。
    * 故Reader会缓冲数据，移动fd的偏移量。如果Reader后不读缓冲却读rd，就会丢失缓冲大小的数据区。
  * 但socket有fd有明显的不同。取决于发送端是否发送数据及网络情况。
    * client先发一个header，再发一个body。前后不是一次write，在网络上可能是两次发送。
    * 所以server端bufio也没有超前读取。

`————以上来自饼干老师的解答`

#### LS file

```shell
### server

# go run miniFtp2.go  -r ../../
2017/08/01 22:30:25 root: ../../
2017/08/01 22:30:25 root: ../../
2017/08/01 22:30:28 cmd:LS, name:day9

# echo 'LS day9' | nc localhost 8021
type	name		size
file	.DS_Store	10244
file	a.txt	14
dir	client	272
file	ftp.txt	10
dir	long	136
dir	miniFtp	204
dir	miniFtp2	272
### ...
```

