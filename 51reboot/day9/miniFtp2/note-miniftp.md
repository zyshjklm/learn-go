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

