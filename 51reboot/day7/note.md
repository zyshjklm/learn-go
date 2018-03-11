

## 作业

重点讲了：

* 如何写map的级联。
* 如何进行序列化。
  * 结构体字段的可见性
  * marshal, unmarshal接口及其方法。


更新数据的2种方式：

```go
if _, ok := c.students[name]; !ok {
        return fmt.Errorf("not exist info for %s", name)
}
c.students[name].id = id

// style 2
if stu, ok := c.students[name]; ok {
        stu.id = id
} else {
        return fmt.Errorf("not exist info for %s", name)
}
```





## 并发

go的并发是语言自身就带有的功能，其他语言通常是基于外部的库，如libevent。

进程与线程

* 进程是最小的资源分配单位，线程是最小调度单位。
* 进程是**管理资源**的
  * 申请内存
  * 打开socket，包括网络，文件等
  * 信号，队列等
* 线程是调度的
  * 早期的linux是没有线程的。
  * 但进程分配资源是一个比较重的任务。比较耗时，比如apache来一个请求分配一次。
  * 与进程是一套调度机制。线程可以分到不同的CPU核上运行。

Java就是线程模型，使用线程池。

C10K问题。每秒10K个并发。



### routine

启动协程：

* go关键字加函数。
* 在正常函数调用前面加上go。

注意：当主函数退出时，所有协程都会退出。当函数退出时，函数里的chan也结束了。

示例：

* routine/sleep.go      


* routine/sort-sleep.go

### channel

协程并不总是独立工作。

需要配合。

#### 示例

**示例1**：将一个切片算总和，使用类似map/reduce的方式计算。

`chan int`是一个整体，表示存放int的channel。本次是子函数存储，主函数读取。

code: channel/sum-int.go

示例2：将一个切片的字符串拼接起来。

code: channel/sum-string1.go

code: channel/sum-string2.go

channel缓冲区

有缓冲后，不需要同步等待。无缓冲时，相当于电话，需要等待对方接电话才能通讯；有缓冲时，就相当于是信箱，可以将多封信放进去，不管对方有没有在。等对方有空了才看。不需要等待对方看。

code: channel/sum-buf.go

#### range channel

range一个channel，有数据则读取，没有则等待。如果close了一个chan，则下次range时会退出。

使用协程来计算fibonacci数。将结果放进channel中。

code:channel/fibo-range.go

#### select channel

上面的例子，由子函数来控制fib数据的生产及结束。主函数比较被动。

通过select多路选择，可能由主函数来主动控制fib的生产。

code:channel/fibo-select.go

注意chan的读写方法：

```go
ch := make(chan int)
quit := make(chan int)
// write
ch <- 3
quit <- 0
// read 
<-ch
<-quit

```



#### 使用channel做计时器

使用time的Ticker对象做计时器。

```go
timer := time.NewTicker(time.Second)
for _ = range timer.C {
    fmt.Println(time.Now().Format("03:04:05"))
}

// 调用 timer.Stop()即可结束定时器。这里只是停止发数据，但没有关闭chan。
timer := time.NewTicker(time.Second)
cnt := 0
for _ = range timer.C {
    cnt++
    fmt.Println(time.Now().Format("03:04:05"))
    if cnt > 10 {
        timer.Stop()
      	return // 通过返回函数，来结束channel。
    }
}
// timer Stop之后，再次range会报错：
// fatal error: all goroutines are asleep - deadlock!
```

code: channel/channel-timer.go



直接使用time的通道功能。

code: channel/channel-boom.go



### 多线程爬虫：抓取图片

基本过程：

* 抓取网页

* 获取图片链接

* 格式化链接

* 下载图片到本地

* 打tar包

  ​

#### http.Get() 

code: http/http-html.go

url的各个部分：

`http://news.baidu.com/ns/cl=2&tn=news#invalidxx`

* **http://**: schema表示协议的类型。如http, https, ftp或者其他的应用层协议。
* **host**: 主机。news.baidu.com
* **路径**：/ns/xxx。
* 请求字符串："?"后面的部分。相当于是变量及其值。
* 锚点：“#”后面的部分。代表当前网面的位置。

url里的字符对于特殊字符会进行转义，包括中文，空格等。

#### 解析链接

code: http/http-uri.go



#### 解析网页

goquery库。类似js中的jquery库。解析网页中的标签。

code: http/http-parse-url.go
通过goquery解析一个网页，得到一个文档，对文档Find某个标签。并直接遍历所有标签，选取其中的某些属性。`doc.Find("img").Each(func(i int, s *goquery.Selection) {})`



#### 格式化链接

`//xx.com/a.jpg`前面的`//`表示和请求的协议相同。因此需要自行拼装。

code: http/http-clean.go

需要注意代码中有两类链接。

* 第一个是请求的原始originURL。比如`http://localhost:8000/http/test.html`
* 第一个是从originURL得到的网页，并从网页中解析出来的图片连接uri.

因此clean()传入的是originURL

repairLINK传入的两个参数：

* uri 是originURL通过url.Parse(originURL)得到的*url.URL变量
* url则是从网页中提取的`img src`中的待清洗链接。

对于相对路径的清洗：

* 先从uri中获取出前置路径，即本示例中`/http/test.html`中的`http`
* 再由前置路径与url（img/b.jpg）进行拼接。

使用本地测试文件test.html

```shell
# ls http
http-clean.go     http-ftp.go       http-html.go      http-parse-url.go http-uri.go       test.html

# nohup python -m SimpleHTTPServer 8000

go run http/http-clean.go 'http://localhost:8000/http/test.html'
-- clean: https://pic1.zhimg.com/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
http://localhost:8000/http/test.html http
-- clean: //pic4.zhimg.com/v2-40becd4a519329198ecb3807f342fd7b.jpg
http://localhost:8000/http/test.html //
-- clean: /golang-spider/img/a.jpg
http://localhost:8000/http/test.html /
-- clean: img/b.jpg
http://localhost:8000/http/test.html relative path: /http/test.html
http://localhost:8000/http/test.html pathS: [ http test.html]
0 https://pic1.zhimg.com/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
1 http://pic4.zhimg.com/v2-40becd4a519329198ecb3807f342fd7b.jpg
2 http://localhost:8000/golang-spider/img/a.jpg
3 http://localhost:8000/http/img/b.jpg

```

以上即可访问：http://localhost:8000/http/test.html

