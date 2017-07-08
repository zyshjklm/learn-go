

### 子函数中的错误处理：

不要在子函数中使用log.Fatal方式，这会直接退出。

这个命令通常用于初始化的时候，比如读取配置，建立连接等不成功。

但程序使用中，用户输入的错误，只需要返回错误给上层，而不是退出。



### shell功能的问题

1 使用io.Pipe后，设置其读和写到相应命令的输入或输出。这里有个问题是没有关闭管道。比如 cat xxx.go | grep test。导致后面的程序一起hung住。

解决办法之一，是cmd.StdoutPipe()函数来代其cmd.Stdout.

另外，cmd.Run()等于cmd.Start() + cmd.Wait(). 这会阻塞程序。而先Start才不会阻塞后面的远行。



文档

* 英文：godoc -http=:6060 &



传参数。同一类的。可以使用args …string，也可以[]string。


