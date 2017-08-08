### tcpProxy2.go

首先需要修正的问题是，在两个go里面，需要关闭的对像不同。

```go
cd tcpProxy2
diff ../tcpProxy/proxy2.go tcpProxy2.go

< 		conn.Close()
---
> 		// 从client读完了，则关闭写server端。
> 		remote.Close()
44c45,46
< 		remote.Close()
---
> 		// 从server端读完了，则关闭写client端。
> 		conn.Close()
```



**测试：使用百度。**

服务端：

```shell
go run tcpProxy2.go
2017/08/08 23:48:20 domain:  www.baidu.com:80
2017/08/08 23:48:22 start to handle conn...

2017/08/08 23:48:22 connect to domain: www.baidu.com:80
2017/08/08 23:48:22 start to wait()...
2017/08/08 23:48:22 go wr start...
2017/08/08 23:48:22 go rd start...
2017/08/08 23:48:22 go wr end...
2017/08/08 23:48:22 go rd end...
2017/08/08 23:48:22 end of wait()...
2017/08/08 23:48:22 end of handle...
2017/08/08 23:48:25 start to handle conn...

2017/08/08 23:48:26 connect to domain: www.baidu.com:80
2017/08/08 23:48:26 start to wait()...
2017/08/08 23:48:26 go wr start...
2017/08/08 23:48:26 go rd start...
2017/08/08 23:48:26 go wr end...
2017/08/08 23:48:26 go rd end...
2017/08/08 23:48:26 end of wait()...
2017/08/08 23:48:26 end of handle...
^Csignal: interrupt
```

client

```shell
curl 127.0.0.1:8021 -H "Host: www.baidu.com"
<!DOCTYPE html>
<!--STATUS OK--><html> <head><meta http-equiv=content-type content=text/html;charset=utf-8><meta http-equiv=X-UA-Compatible content=IE=Edge><meta content=always name=referrer><link rel=stylesheet type=text/css href=http://s1.bdstatic.com/r/www/cache/bdorz/baidu.min.css><title>百度一下，你就知道</title></head> <body link=#0000cc> <div id=wrapper> <div id=head> <div class=head_wrapper> <div class=s_form> <div ....

curl -v 127.0.0.1:8021
* Rebuilt URL to: 127.0.0.1:8021/
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8021 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8021
> User-Agent: curl/7.49.1
> Accept: */*
>
* Empty reply from server
* Connection #0 to host 127.0.0.1 left intact
curl: (52) Empty reply from server
```

不带Host时，使用了Ip。因此请求出错。



测试：使用qq

```shell
#server

go run tcpProxy2.go --domain=www.qq.com:80
2017/08/08 23:55:40 domain:  www.qq.com:80
2017/08/08 23:55:42 start to handle conn...

2017/08/08 23:55:42 connect to domain: www.qq.com:80
2017/08/08 23:55:42 start to wait()...
2017/08/08 23:55:42 go wr start...
2017/08/08 23:55:42 go rd start...
2017/08/08 23:55:42 go rd end...
2017/08/08 23:55:42 go wr end...
2017/08/08 23:55:42 end of wait()...
2017/08/08 23:55:42 end of handle...

### client
curl -v 127.0.0.1:8021
* Rebuilt URL to: 127.0.0.1:8021/
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8021 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8021
> User-Agent: curl/7.49.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: squid/3.5.20
< Date: Tue, 08 Aug 2017 15:55:42 GMT
< Content-Type: text/html
< Transfer-Encoding: chunked
< Connection: keep-alive
< Vary: Accept-Encoding
< Expires: Tue, 08 Aug 2017 15:55:41 GMT
< Cache-Control: no-cache
< X-Cache:  from tianjin.qq.com
<
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<meta http-evuiv="Access-Control-Allow-Origin" content="*">
<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum ...

```

即使Host: 127.0.0.1，也能返回。qq对访问没有这么严格。



**测试：ssh**

```shell
### server

go run tcpProxy2.go --domain=127.0.0.1:22
2017/08/09 00:00:56 domain:  127.0.0.1:22
2017/08/09 00:01:21 start to handle conn...

2017/08/09 00:01:21 connect to domain: 127.0.0.1:22
2017/08/09 00:01:21 start to wait()...
2017/08/09 00:01:21 go wr start...
2017/08/09 00:01:21 go rd start...
2017/08/09 00:01:45 go rd end...
2017/08/09 00:01:45 go wr end...
2017/08/09 00:01:45 end of wait()...
2017/08/09 00:01:45 end of handle...

### client

ssh -p 8021 song@127.0.0.1
The authenticity of host '[127.0.0.1]:8021 ([127.0.0.1]:8021)' can't be established.
ECDSA key fingerprint is SHA256:arfUZ081KHzVw0e4YVQ41fIW3yVXHyFxmdNqC64sXkg.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '[127.0.0.1]:8021' (ECDSA) to the list of known hosts.
Password:
Last login: Tue Aug  8 10:41:43 2017
➜  ~ ps
### ...

➜  ~ exit
Connection to 127.0.0.1 closed.
➜  tcpProxy git:(master) ✗


```



**访问：星球大战**

```shell
### server
go run tcpProxy2.go --domain=towel.blinkenlights.nl:23
2017/08/09 00:05:03 domain:  towel.blinkenlights.nl:23
2017/08/09 00:05:42 start to handle conn...

2017/08/09 00:05:43 connect to domain: towel.blinkenlights.nl:23
2017/08/09 00:05:43 start to wait()...
2017/08/09 00:05:43 go wr start...
2017/08/09 00:05:43 go rd start...

### client

telnet 127.0.0.1 8021
### 观看 星球大战。

```



#### 简写方式

```shell
### server 45行左右
go run short.go
2017/08/09 00:15:59 domain:  www.qq.com:80
2017/08/09 00:16:03 start to handle conn...
2017/08/09 00:16:11 start to handle conn...

### client

curl -v 127.0.0.1:8021
* Rebuilt URL to: 127.0.0.1:8021/
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8021 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:8021
> User-Agent: curl/7.49.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: squid/3.5.20
< Date: Tue, 08 Aug 2017 16:16:11 GMT
< Content-Type: text/html
< Transfer-Encoding: chunked
< Connection: keep-alive
< Vary: Accept-Encoding
< Expires: Tue, 08 Aug 2017 16:16:10 GMT
< Cache-Control: no-cache
< X-Cache:  from tianjin.qq.com
<
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<meta http-evuiv="Access-Control-Allow-Origin" content="*">
curl -v 127.0.0.1:8021

```

