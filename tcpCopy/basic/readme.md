
refer: 《用 Go 实现 TCP 连接的双向拷贝》

## basic 

使用多个终端进行观察

##### 1 建立连接

```shell
# T1 运行另一台虚机
nc -l 8849

# T2 
go run basic/main.go

# T3 listen
netstat -an | grep '884[89]'

tcp4       0      0  127.0.0.1.8848     *.*               LISTEN
```

##### 2 请求数据

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

nc -l 8849

from 8848: hello golang
from 8849: HELLO GOLANG

```

##### 3 观察连接

```shell
netstat -an | grep '884[89]'
tcp4  192.168.100.106.57516  192.168.100.212.8849   ESTABLISHED
# 本地8848启动随机端口57516，连接远端8849

tcp4  127.0.0.1.8848         127.0.0.1.57515        ESTABLISHED
tcp4  127.0.0.1.57515        127.0.0.1.8848         ESTABLISHED
# 本地telnet启动57515连接8848

tcp4  127.0.0.1.8848         *.*                    LISTEN
```

最后中断任意一端。双端都退出。直接TIME_WAIT。

* 退本地telnet，则如下图。
* 退远端8849。则本地只能看到8848连向telnet的随机端口。

```shell

netstat -an | grep '127.0.0.1.884[89]'
tcp4       127.0.0.1.8848         *.*                    LISTEN
tcp4       127.0.0.1.57515        127.0.0.1.8848         TIME_WAIT
tcp4       192.168.100.106.57516  192.168.100.212.8849   TIME_WAIT

```

