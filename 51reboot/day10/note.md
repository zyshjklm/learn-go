

## 第10课 代理



### 代理服务器

**正向代理服务**

C要访问S。因为访问限制，不能直接访问。比如不能直接访问google。现使用代理程序p。

**( client -> proxy ) -> server**。s响应返回给p，p再将响应返回给c。

正向代理是**针对client端**的。

**反向代理服务**

服务端要提供访问。但为了安全，不将server直接暴露在公网上。而是将代理放在公网上，将服务放在代理的后端，通常在内网。原本的c -> s模式 变成了：

**client -> ( proxy  -> server)**。

对client来说，只能看到proxy。由其响应请求。

反向代理是针对于**server端**。



### 实现tcp层代理

tcpProxy/tcpProxy.go

#### 1 访问网页

server端

```shell
go run tcpProxy/tcpProxy.go --domain 'www.qq.com:80'
2017/08/06 17:10:10 domain:  www.qq.com:80
2017/08/06 17:10:12 start to handle conn...

2017/08/06 17:10:12 connect to domain: www.qq.com:80
2017/08/06 17:10:12 start to wait()...
2017/08/06 17:10:12 go wr start...
2017/08/06 17:10:12 go rd start...
2017/08/06 17:10:12 go rd end...
2017/08/06 17:10:13 start to handle conn...

2017/08/06 17:10:13 connect to domain: www.qq.com:80
2017/08/06 17:10:13 start to wait()...
2017/08/06 17:10:13 go wr start...
2017/08/06 17:10:13 go rd start...
2017/08/06 17:10:13 go rd end...
2017/08/06 17:10:29 start to handle conn...


go run tcpProxy/tcpProxy.go  --target=www.qq.com:80

curl -v 127.0.0.1:8021


```

client

```shell
curl  127.0.0.1:8021 >qq1.html
curl  127.0.0.1:8021 >qq2.html
ls -l qq1.html qq2.html
-rw-r--r--  1 song  staff  34379 Aug  6 17:13 qq1.html
-rw-r--r--  1 song  staff  34379 Aug  6 17:13 qq2.html
```

但是使用baidu, 163等网站却会导致curl端到8021的连接未关闭，应该是将第一个协程中的io.Copy(remote, conn)未正常结束。

#### 2 ssh访问

```shell
# server 
go run tcpProxy/tcpProxy.go --domain localhost:22
2017/08/06 17:16:21 domain:  localhost:22
2017/08/06 17:16:35 start to handle conn...

2017/08/06 17:16:35 connect to domain: localhost:22
2017/08/06 17:16:35 start to wait()...
2017/08/06 17:16:35 go wr start...
2017/08/06 17:16:35 go rd start...
2017/08/06 17:16:48 go wr end...
2017/08/06 17:17:00 go rd end...
2017/08/06 17:17:00 end of wait()...
2017/08/06 17:17:00 end of handle...
2017/08/06 17:17:02 start to handle conn...

## client
nc  127.0.0.1 8021
SSH-2.0-OpenSSH_7.2
ssh song@localhost
Protocol mismatch.
^C
```

看上去是使用方法不对。需要再研究一下。

