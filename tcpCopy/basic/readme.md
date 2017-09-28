
refer: 《用 Go 实现 TCP 连接的双向拷贝》

## basic 

使用多个终端进行观察

建立连接

```shell
# T1
# nc -l 8849
nc 127.0.0.1 -l 8849

# T2 
go run basic/main.go

# T3 listen
netstat -an | grep '127.0.0.1.884[89]'

tcp4       0      0  127.0.0.1.8849         *.*                    LISTEN
tcp4       0      0  127.0.0.1.8848         *.*                    LISTEN

```

请求数据

```shell

# T4 连接，并发送数据，观察8849的输出，再从8849输入并回车。

telnet localhost 8848
Trying ::1...
telnet: connect to address ::1: Connection refused
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.

from 8848: hello golang
from 8849: HELLO GOLANG

# T1

nc 127.0.0.1 -l 8849

from 8848: hello golang
from 8849: HELLO GOLANG


```

最后中断任意一端。双端都退出。

