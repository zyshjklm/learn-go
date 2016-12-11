## gopl chapter 5

#### How to add a package from Internet

for example, chapter 5 need "golang.org/x/net/html".

because of the GFW, you need git clone the src code.

```shell

$GOPATH

mkdir -p $GOPATH/src/golang.org/x/
cd $GOPATH/src/golang.org/x/

git clone https://github.com/golang/net

# now, you can use "golang.org/x/net/html"

## update, just use 
go get https://github.com/golang/net

```

#### findlink

```shell

go build ../ch1/fetch/fetchall.go
# fetchall

go build findlinks1/main.go

./fetchall https://golang.org > go.html

cat go.html | ./main

```

print all the node type and data

```shell

go build findlinks/main.go

cat go.html | ./main
Type: 2, Data:
Type: 5, Data: html
Type: 3, Data: html
Type: 3, Data: head
Type: 1, Data:
Type: 3, Data: meta
Type: 1, Data:
Type: 3, Data: meta
Type: 1, Data:
Type: 3, Data: meta
...

```

#### wait and retries

```shell

2016/12/11 17:59:08 Head url http://golang.org
2016/12/11 17:59:15 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
2016/12/11 17:59:16 Head url http://golang.org
2016/12/11 17:59:18 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
2016/12/11 17:59:20 Head url http://golang.org
2016/12/11 17:59:22 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
2016/12/11 17:59:26 Head url http://golang.org
2016/12/11 17:59:28 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
2016/12/11 17:59:36 Head url http://golang.org
2016/12/11 17:59:38 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
2016/12/11 17:59:54 Head url http://golang.org
2016/12/11 17:59:57 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
Site is down: server http://golang.org failed to respond after 1m0s
exit status 1

```

每一次Head url的时间比前一行的时间差是指数增加的：
* 0: 08-08
* 1: 16-15
* 2: 20-18
* 4: 26-22
* 8: 36-28
* 16: 54-38
* 32:


修改前缀
```shell

go run wait/main.go http://golang.org
[wait] 2016/12/11 18:21:10 Head url http://golang.org
[wait] 2016/12/11 18:21:17 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
[wait] 2016/12/11 18:21:18 Head url http://golang.org
[wait] 2016/12/11 18:21:20 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
[wait] 2016/12/11 18:21:22 Head url http://golang.org
[wait] 2016/12/11 18:21:24 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
[wait] 2016/12/11 18:21:28 Head url http://golang.org
[wait] 2016/12/11 18:21:30 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
[wait] 2016/12/11 18:21:38 Head url http://golang.org
[wait] 2016/12/11 18:21:41 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
[wait] 2016/12/11 18:21:57 Head url http://golang.org
[wait] 2016/12/11 18:22:00 server not responding (Head http://golang.org: dial tcp 216.239.37.1:80: i/o timeout); retrying...
2016/12/11 18:22:32 Site is down: server http://golang.org failed to respond after 1m0s
exit status 1

```

#### 错误处理的5种策略

* 向上层传递。直接传递或者增加前缀
* 做有意思的重试，常用指数回退法控制重试的时间
* 打印错误并优雅退出
* 打印错误并继续执行
* 忽略错误，仅某些特殊情况下使用。比如删除某个日志文件


