## chat server

每个用户使用如下任意方式连接：

* nc localhost 8021
* telnet localhost 8021



client

```shell
# telnet localhost 8021
Trying ::1...
Connected to localhost.
Escape character is '^]'.
	pls input you name and pwd
jungle 123456
	login success
hello
song: hello all
who are you
```

client

```shell
# nc localhost 8021
	pls input you name and pwd
junsdfg 2134
	bad password
	pls input you name and pwd
song 123456
	login success
hello all
jungle: who are you
```



server

```shell
# go run chatServer.go
2017/08/02 00:42:57 [jungle] loggined
2017/08/02 00:42:57 usr:jungle, pwd:123456
2017/08/02 00:42:57 [worker] [jungle] start to check
2017/08/02 00:42:57 [worker] [jungle] waiting msg...
2017/08/02 00:43:24 [broadcase] user num:1
2017/08/02 00:43:24 [broadcase] user: jungle	msg:hello
2017/08/02 00:43:24 [worker] [jungle] start to check
2017/08/02 00:43:24 [worker] [jungle] waiting msg...
2017/08/02 00:43:40 [song] loggined
2017/08/02 00:43:40 usr:song, pwd:123456
2017/08/02 00:43:40 [worker] [song] start to check
2017/08/02 00:43:40 [worker] [song] waiting msg...
2017/08/02 00:43:42 [broadcase] user num:2
2017/08/02 00:43:42 [broadcase] user: jungle	msg:hello all
2017/08/02 00:43:42 [broadcase] user: song	msg:hello all
2017/08/02 00:43:42 [worker] [song] start to check
2017/08/02 00:43:42 [worker] [song] waiting msg...
2017/08/02 00:43:57 [golang] loggined
2017/08/02 00:43:57 usr:golang, pwd:123456
2017/08/02 00:43:57 [worker] [golang] start to check
2017/08/02 00:43:57 [worker] [golang] waiting msg...
2017/08/02 00:44:05 [broadcase] user num:3
2017/08/02 00:44:05 [broadcase] user: golang	msg:hello i am golang
2017/08/02 00:44:05 [broadcase] user: jungle	msg:hello i am golang
2017/08/02 00:44:05 [broadcase] user: song	msg:hello i am golang
```

