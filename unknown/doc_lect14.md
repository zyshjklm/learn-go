## lect14  并发concurrency

### goroutine

* goroutine只是由官方实现的超级“线程池”而已，但每个实例4-5K的栈内存和大幅减少的创建与销毁开销，是其并发的根本原因。


* 并发不是并行。concurrency in not parallelism


* 并发主要是由切换时间片来实现“同时”运行，并行则是直接利用多核实现多线程的运行


* go可以设置使用核数，以发挥多核能力


* goroutine奉行**通过通信来共享内存**，而不是共享内存来通信。



### channel

* channel是goroutine沟通的桥梁，大都是阻塞同步的
* 通过make创建，close关闭
* channel是引用类型
* 可以通过for range来迭代操作channel
* 可以设置单向或双向通道。默认make是双向。
* 可以设置缓存大小，在未被填满前不会发生阻塞



### select

* 可处理一个或多个channel的发送与接收
* 同时有多个可用的channel时按随机顺序处理
* 可用空的select来阻塞main函数
* 可设置超时

