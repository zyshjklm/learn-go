

### crypto V3

V2版本的错误，已经在其[目录下说明](../cryptoSocks5-V2/crypto-v2.md)。

正确的版本，其思路应该是[V1版本](../cryptoSocks5-V1/cryptoSocks5.md)中的说明来实现。其实核心代码就2行。

![conn-remote-encrypt](../cryptoSocks5-V1/conn-remote.png)

这里再简述一下：

加密（以上图为加密端）：

* 对remote端进行Reader/Writer封装

解密（以上图为解密端）：

* 对conn端进行Reader/Writer封装



因此，就像在打包基础上压缩一样，对加密的数据进行解密也是一样的。原本是直接读写conn，现在变成对conn进行一层封装。

错误开始于：

```go
// 原本实现是这样的
r := bufio.NewReader(client)
addr, err := socks5Auth(r, client)

// 封装client连接
clientRd := mycrypto.NewCryptoReader(conn, key)
clientWr := mycrypto.NewCryptoWriter(conn, key)

rdBuf := bufio.NewReader(clientRd)
addr, err := socks5Auth(rdBuf, clientWr)

```

2017.08.17.



以下是讲解：

### 进程通信方式

信号量

sockets

共享内存

### proxy_b的功能

- 将数据解密
- 转发给C
- 加密C的数据



怎么去掉转发——直接将数据转发给C——通过封装进行转发。

在原来的C基础上进行封装，进行解密。

其实就是将B去掉，由C直接负责对conn进行解密。
——将一个reader进行解密。

**crypto的作用**：将一个reader/writer封装成另一个可以加密/解密的reader/writer。

### 面向接口的编程

在传递参数时，要传小接口，而不是大接口。
比如：
`func NewCryptoReader(r io.Reader, key string) io.Reader {}`
上面的r只是需要一个io.Reader，你给一个net.Conn也是可以的。但这样就不能用于基于文件的Reader了。



