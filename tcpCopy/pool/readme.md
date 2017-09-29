
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
tcp4  192.168.100.106.57657  192.168.100.212.8849   ESTABLISHED
# 本地8848启动随机端口57516，连接远端8849

tcp4  127.0.0.1.8848         127.0.0.1.57656        ESTABLISHED
tcp4  127.0.0.1.57656        127.0.0.1.8848         ESTABLISHED
# 本地telnet启动57656连接8848
tcp4  127.0.0.1.8848         *.*                    LISTEN
```

##### 4 最后测试退出。

* 退出远端

```shell

### 退出远端，
netstat -an | grep '127.0.0.1.884[89]'
tcp4  192.168.100.106.57657  192.168.100.212.8849   CLOSE_WAIT
tcp4  127.0.0.1.8848         *.*                    LISTEN
tcp4  127.0.0.1.8848         127.0.0.1.57656        TIME_WAIT
# 57656等一会会消失。57657会一直存在。

# 再次启动远端8849，并在本地发起连接，直接失败退出。
telnet localhost 8848
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Connection closed by foreign host.

```

因为**57657端口实际上已经不可用了。但却没有关闭**。

* 重启程序和连接。

```shell
netstat -an | grep '884[89]'
tcp4  127.0.0.1.8848         127.0.0.1.57727        ESTABLISHED
tcp4  127.0.0.1.57727        127.0.0.1.8848         ESTABLISHED
tcp4  192.168.100.106.57724  192.168.100.212.8849   ESTABLISHED
tcp4  127.0.0.1.8848         *.*                    LISTEN

# 断开telnet，telnet启动的随机端口57727进入TIME_WAIT，随后会消失
tcp4  127.0.0.1.57727        127.0.0.1.8848         TIME_WAIT
tcp4  192.168.100.106.57724  192.168.100.212.8849   ESTABLISHED

# 再次启动telnet时，可以正常发送数据。因为8848到8849端并不知道telnet的变化
tcp4  127.0.0.1.8848         127.0.0.1.57750        ESTABLISHED
tcp4  127.0.0.1.57750        127.0.0.1.8848         ESTABLISHED
tcp4  192.168.100.106.57724  192.168.100.212.8849   ESTABLISHED
```

继续使用57724转发到8849端。

