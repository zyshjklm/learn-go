## day 11



### 加密隧道

通讯的两端都在TCP层进行加解密和socks5代理。

数据流向：

app --- agent/encrypt — (-- network-- ) -- proxy/decrypt -- socks5  --- appServer

* A agent
* B proxy
* C socks5

proxy和socks5是可以合在一起的。在设计上先分开实现。

之前实现的day10/tcpProxy2/tcpProxy2.go，即可用来完成A和B位置的工作。只是需要补充些功能：

* A端：监听请求并转发数据到目标端。
  * 初略说明：
    * 对conn请求要发送的数据加密。
    * 对remote返回的数据进行解密。
  * 准确说明：
    * 对conn端的数据不做处理
    * 对remote要写入的数据进行加密；对remote中读取的数据进行解密
* B端：监听远端请求并转给代理端
  * B端和A端是隧道的两端，其操作是相对的，加解密都在conn端。

### 如何加解密

参考rc4.go的实现方式，以及tar包的压缩。通过接口进行封装和过滤。

对于A端，根据说明，只需要对remote进行封装：

* 写：封装remote为Writer
  * 在Write函数中对原始buf中的byte数据进行加密，再将其写到remote。
* 读：封装remote为Reader
  * 在read函数中通过remote的Read先读到buf中，再进行一次解密操作

其背后的原理，就是interface。一个结构体，实现了接口的所有方法，即实现了该接口。remote发送数据时是个Writer，因此封装它时，也需要实现一个Writer。



### 测试

将源代码放在一个目录下，文件内packge即可以是main，也可以是其他值，如mycrypt。再增加一个xxx_test.go的文件。用来写单测。

测试代码的package要与源代码相同。单测函数名是TestXXX样式。

测试时并没有使用真实了网络传输，而是使用内存文件bytes.Buffer。



具体代码参考：

../cryptoSocks5-V1/mycrypto/mycrypto.go

其中main函数中的2段代码，分别从stdin, stdout这两个角度实现的。



后继说明参见../cryptoSocks5-V1/cryptoSocks5.md

