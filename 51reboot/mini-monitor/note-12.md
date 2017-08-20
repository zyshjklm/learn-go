## mini-monitor

第12次课。第一次监控课。



先在testAgent/目录下写了三轮小的采集agent。下面开始统一在一个目录下进行迭代完善。

#### 1）第一轮 使用sender来进行发送。

* agent/main.go 基于testAgent/agent3.go文件
* agent/sender.go，发送端。这样在主程序中，不用关心发送的细节。

需要注意的事，两个文件都是`package main`，这在编译时需要注意。

```shell
# go run main.go
./main.go:33: undefined: NewSender
#### 编译失败，因为函数在sender.go文件中


# go build
# ls -l agent
-rwxr-xr-x  1 song  staff  3756736 Aug 20 14:59 agent

###  另一个终端先启动端口监听，再启动agent
# nc -l 127.0.0.1 6000
{"metric":"cpu.usage","endpoint":"jungle85","tag":["darwin"],"value":15.384615384615385,"timestamp":1503212470}
{"metric":"mem.usage","endpoint":"jungle85","tag":["darwin"],"value":526478.90625,"timestamp":1503212470}

# ./agent

```

也可以直接编译多个文件

```shell
###  另一个终端先启动端口监听，再启动agent
# nc -l 127.0.0.1 6000

# go run sender.go main.go
```

这样6000端口也能收到数据。

#### 2）第二轮 完善sender中的远端连接与重试。

agent/sender.go完善：

- 新增connect()
  - 建立到远端transfer的连接，失败时按2的指数延迟重试。
  - 最小间隔0.1秒，最大间隔30秒。
- 新增reConnect()，当发送中失败时，用于清理conn并调用connect重建连接
- 完善Start() ：
  - 使用bufio封装Writer。
  - 当有数据到达时，将数据写入Writer
  - 基于ticker，每5秒刷新写一次Writer。而不是每来一条记录就写一次

尝试运行：

```shell
# nc -l 127.0.0.1 6000

# go run main.go sender.go
2017/08/20 15:35:49 local addr:127.0.0.1:55327
2017/08/20 15:36:24 Flush to remote err:write tcp 127.0.0.1:55327->127.0.0.1:6000: write: broken pipe
2017/08/20 15:36:24 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:24 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:24 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:25 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:26 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:27 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:31 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:37 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:36:50 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:37:15 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:37:45 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:38:15 dial tcp :6000: getsockopt: connection refused
2017/08/20 15:38:45 dial tcp :6000: getsockopt: connection refused
```

启动监听后，每5秒收到来自agent发送的数据。如果关闭监听端。则会看到agent的连接重试过程。前面几次间隔比较小，区别不明显。之后从4，8，16，30则比较明显看出来。

#### 3）第三轮 增加sched和builtin。

完善：

- builtin.go 用于存储各种内建的监控项，如cpu/mem/disk等
- sched.go 用于调度所有的监控项。
- main.go 
  - 主函数只需要通过sched添加监控metric即可。
  - 增加debug开头。

```shell
# go build

# ./agent
2017/08/20 22:52:02 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:52:02 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:52:02 dial tcp :6000: getsockopt: connection refused

# nc -l 127.0.0.1 6000
{"metric":"mem.percent","endpoint":"jungle85","tag":["darwin"],"value":554504601600,"timestamp":1503240725}
{"metric":"mem.used","endpoint":"jungle85","tag":["darwin"],"value":5545046016,"timestamp":1503240725}
{"metric":"mem.percent","endpoint":"jungle85","tag":["darwin"],"value":554321100800,"timestamp":1503240728}
{"metric":"mem.used","endpoint":"jungle85","tag":["darwin"],"value":5543211008,"timestamp":1503240728}
{"metric":"cpu.usage","endpoint":"jungle85","tag":["darwin"],"value":3.482587064676617,"timestamp":1503240728}
{"metric":"cpu.load1","endpoint":"jungle85","tag":["darwin"],"value":1.22,"timestamp":1503240728}
{"metric":"cpu.load5","endpoint":"jungle85","tag":["darwin"],"value":1.44,"timestamp":1503240728}

```

debug

```shell
# nc -l 127.0.0.1 6000

### ignore the output

# ./agent -h
Usage of ./agent:
  -debug
    	debug data output
  -trans string
    	transfer address (default ":6000")

#./agent -debug

2017/08/20 22:56:28 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:56:28 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:56:28 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:56:28 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:56:29 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:56:31 ## ticker hit at2017-08-20 22:56:31.140373492 +0800 CST
2017/08/20 22:56:31 &{mem.percent jungle85 [darwin] 5.583228928e+11 1503240991}
2017/08/20 22:56:31 &{mem.used jungle85 [darwin] 5.583228928e+09 1503240991}
2017/08/20 22:56:31 dial tcp :6000: getsockopt: connection refused
2017/08/20 22:56:33 ## ticker hit at2017-08-20 22:56:33.145445037 +0800 CST
2017/08/20 22:56:34 ## ticker hit at2017-08-20 22:56:34.140461537 +0800 CST
2017/08/20 22:56:34 &{mem.percent jungle85 [darwin] 5.58413824e+11 1503240994}
2017/08/20 22:56:34 &{mem.used jungle85 [darwin] 5.58413824e+09 1503240994}
2017/08/20 22:56:34 &{cpu.usage jungle85 [darwin] 9.476309226932669 1503240994}
2017/08/20 22:56:34 &{cpu.load1 jungle85 [darwin] 1.47 1503240994}
2017/08/20 22:56:34 &{cpu.load5 jungle85 [darwin] 1.52 1503240994}
2017/08/20 22:56:34 local addr:127.0.0.1:49473
2017/08/20 22:56:34 ~~ sender get metric
2017/08/20 22:56:34 ~~ sender get metric

```

