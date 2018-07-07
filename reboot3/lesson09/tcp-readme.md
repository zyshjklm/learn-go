## tcp test



```shell
# mkdir tcptest
# cd tcptest

# cobra init
# cobra add client
# cobra add server
# go run main.go

```



### add pkg

```shell
# cd cmd
# mkdir common tcpClient tcpServer

# vim common/common.go
# vim tcpClient/tcpClient.go
# vim tcpServer/tcpServer.go

# vim client.go
# vim server.go

# cd ../
```

server

```shell
# go run main.go server
server called
listenning...
conn from 127.0.0.1:58538
conn from 127.0.0.1:58543

```

client

```shell
# go run main.go client
client called
send...
recv...
[2018-07-07 12:47:40.196412526 +0800 CST m=+0.002932981] this is a client ...

```

