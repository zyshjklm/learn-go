

## homework



### 示范使用一个chan和一个协程

* 先写一个printURL函数。给定url，抓取url并打印状态

* 定义一个work函数，接收ch中来的url。因此参数是work(ch chan string)

  * work调用printURL进行处理

* 主协程主要工作：

  * 创建chan，
  * 启动协程
  * 向chan写入数据

* work获取chan并处理

  ​


### 使用一个chan及多个协程

* 主协程：
  * 建一个chan
  * 启动三个协程work()
  * 多次向协程中写入url
  * 关闭chan
  * sleep等待
* work协程
  * 循环从chan中获取url【可使用for range chan】
  * 调用printURL()
  * 直到chan关闭



### channel的特性：

* 1）只要chan没有关闭。可以源源不断发送数据和接受数据
* 2）如果chan中没有数据，接收方会阻塞
* 3）如果没有人正在等待chan的数据，发送方会阻塞
* 4）从一个已经close掉的chan中读取数据，启动不会阻塞，且得到一个默认值。
* 5）一个chan可以被多个routine竞争数据，有随机性。调度器尽量保持公平
* url, ok := <- ch；当ok为false时，chan已经关闭。

永远不要使用time.Sleep来进行协程同步。一定会出问题。



- ​

### 使用一个chan及多个协程加WaitGroup

- 主协程：
  - 建一个chan
  - 创建一个WaitGroup，并调用WaitGroup.Add
  - 启动三个协程work()
  - 多次向协程中写入url
  - 关闭chan
  - 调用WaitGroup.Wait等待协程结束
- work协程
  - 循环从chan中获取url【可使用for range chan】
  - 调用printURL()
  - 直到chan关闭
  - 循环完成后执行WaitGroup.Done()