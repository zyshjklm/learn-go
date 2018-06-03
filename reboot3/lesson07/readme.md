
## lesson07

### 一、相关坑

1）字符串不能直接修改。

```shell
# 直接修改会报错
./string.go:6:7: cannot assign to s[2]
```

修改需要另行转换。
```shell
# go run string.go
hewlo golang
```
hello.go字符串的长度，以及类型转换。



2）结构体操作

不能与普通变量混合赋值。

struct1.go



3）slice的初始化

需要注意make时缓冲区的大小。

slice1.go



4）defer

注意defer的执行顺序。以及使用panic与recover的执行顺序。

```shell
# go run defer1.go
start call...
start panic...
3
2
1
in defer
recover!
```

调用call函数。输出2行print，进入panic，call函数结束。

然后开始按定义的逆序进行defer的执行，故输出3，2，1.

再回到main函数。结束时执行defer。

defer2.go

```shell
# go run defer1.go
打印panic的调用栈。
```

关键是使用runtime的debug.PrintStack()函数。



5）iota相关操作

iota1.go

iota2.go

iota3.go

iota-byte.go

c语言的const是常量吗，更应该理解成只读变量。



6）for循环操作struct

slice-struct的赋值与循环

* foreach1.go
  * 示范创建结构体的slice，以及赋值。
* foreach2.go
  * 循环slice结构体，并赋值给map
* foreach3.go
  * 修改部分结构体。

预测结构体需要使用的长度与容量，预先分配。避免过多的动态分配。出现性能抖动。



7）heap

```shell
# go tool compile -S heap1.go | more

# go tool compile -m heap1.go

```



8）func函数值

func1.go





### 二、goroutine

* routine0.go
  * 主程序太快，worker还等不等执行就退出了。最笨的办法是使用sleep解决。
* routing1.go
  * 一般使用WaitGroup来同步。
* routing2-timesort.go
  * 通过sleep时间来完成整数的排序。并使用sync.WaitGroup同步。
* routing2-chansort.go
  * 通过通道来保存排序的结果。

通道同步

* routine3-lock.go
  * 示范一个错误的使用。会出现死锁。
  * 因为ch不关闭，匿名函数的wg一起不结束。
* routing3-dead.go
  * 通过for-select模式，持续进行chan的读取操作。
  * 如果没有chan变化，则执行default操作，并等待50毫秒
  * 如果有多个chan变化，则随机选择一个执行。
  * 因为时间设置为2：1，故每一次select到结果之前会执行2次default
  * 当5个worker都返回时。
    * wg的等待结束。
    * 并执行对quit的写操作，让for-select循环return而结束。
* routine4-sum.go
  * 通过两个协程，将slice分成2半进行求和，再对结果二次求和。



chanbuf1.go，设置chan的缓冲长度。

fib1.go，通过2种方式计算fib函数的值。

