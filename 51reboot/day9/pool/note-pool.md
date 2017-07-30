### pool



* **版本1** 
  * 使用协程和单个channel
  * 每个协程只处理一次请求
  * 使用Sleep来等待完成

```shell
# go run pool/pool1.go
http://www.baidu.com 200 OK

```


* **版本2** 
  * 使用多个协程和单个channel
  * 每个协程只处理一次请求
  * 使用Sleep来等待完成

```shell
# go run pool/pool2.go
http://www.baidu.com 200 OK
http://qq.com 200 OK
http://daily.zhihu.com 200 OK
```



* **版本3——大修改**  
  * 使用多个协程和单个channel
  * 每个协程每次循环能处理一次请求
  * 使用Sleep来等待完成

```shell
# go run pool/pool3.go
http://www.baidu.com 200 OK
http://qq.com 200 OK
http://daily.zhihu.com 200 OK
```

如何让协程能不断的处理请求

```go
func work(ch chan string) {
	for {
		url, ok := <-ch
		if !ok {
			break
		}
		printURLStatus(url)
	}
	return
}
```

channel 的几个特性：

* 只要不被close，可以永远发送数据和接受数据
* 如果channel中没有数据，接收方会阻塞
* 如果没有人正在处理channel的数据，发送方会阻塞。
* 因此创建好channel后不能马上向其写数据，而是要先启用协程并使用channel，如 go work(ch)，这样才会有消费方
* close掉channel时，所有使用它的协程在取数据时得到默认值。
* channel的长度是动态的。不能当成是数组。
* 本次循环可能没有数据，则会导致channel阻塞。直到收到有数据，则唤醒本次循环。



- **版本4——range channel** 
  - 使用多个协程和单个channel
  - 每个协程使用**range来遍历channel**
  - 使用Sleep来等待完成

```shell
# go run pool/pool_range.go
http://www.baidu.com 200 OK
http://qq.com 200 OK
http://daily.zhihu.com 200 OK
```

如何让协程能不断的处理请求，下面的方式更简洁优雅：

```go
func work(ch chan string) {
	for url := range ch {
		printURLStatus(url)
	}
	return
}
```

说明：

* range + channel。channel是动态的。而range并不知道其长度。
* 但根据当前状态进行取值，没有则阻塞。
* 协程数量与channel数量：
  * 协程比channel多。一个channel，可以被多个协程竞争抢到
  * 就像一块骨头被几只狗抢。
  * 调度器会尽量保证调度的公平性。调度的算法及其实现
  * 这是生产-消费模式
  * 也像用户请求到达网站入口，被调度到不同后端机器





- **版本5——waitGroup** 
  - 使用多个协程和单个channel
  - 每个协程使用range来遍历channel
  - 使用waitGroup来等待完成，而不是Sleep

```shell
# go run pool/pool_wg.go
http://www.baidu.com 200 OK
http://qq.com 200 OK
http://daily.zhihu.com 200 OK
```

waitGroup方式：

* wg内部是计数器，Add加数，Done减数。Wait在大于0时阻塞。<=0时结束。
* wg.Add(urlNum)
  * Add应该在启动协程前启动。
  * 添加wg的2种方式。
    * 一种是一次添加协和数量个数Add(urlNum)；
    * 一种是每启动一个协程时Add(1)。
* defer wg.Done()的完成是在协程中表示当前协程的结束。

