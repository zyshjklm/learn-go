## 作业

### **点评问题**

前2个问题的修正，反映到51reboot/cryptoSocks5-V1/socks5/中。

#### 1）代理程序

```go
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
	}()
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
	}()
```

这里需要使用对conn的bufio。而不是直接用原始变量。在良好的网络通讯连接中，使用conn看不出啥问题，但当网络质量复杂且延时大时，就可能出现异常。

#### 2）socks5错误处理

代理程序有不少地方都需要做错误处理。可以统一进行。

```Go
func mustReadByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

// readAddr
```

惯例：使用must开始命令的函数，一般在有错误时都会发生panic的。

这种方式主要用于解析协议，有任何错误都直接退出。





#### 2）ftp 使用LS命令时，如果文件比较多，只得到了部分。因为

```go
n,err := conn.Read(buf)
// 只读了一次buf，没有做到循环按块处理。
```



#### 3）类型断言。

只能对接口进行断言。

51reboot/cryptoSocks5-V1/socks5/socks5.go中，因为使用了mustReadByte函数，需要对recover()的返回值进行断言。

```go
defer func() {
    e := recover() // interface{}
    if e != nil {
        err = e.(error)
    }
}()
```

e是一个空的interface{}，将其断言为为错误，如果成功，则err是错误类型，如果断言失败。则继续panic。


