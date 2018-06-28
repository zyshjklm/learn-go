
2018 0623

## 并行并发

进程：

*   内存空间
    *   代码段
    *   数据段
*   句柄
    *   文件
    *   设备
*   线程
    *   初始线程被称为主线程
    *   子线程（CPU）
    *   代码（CPU）



### routine&GOMAXPROCS

例子：goroutine输出字母表

case1: 设置runtime的cpu核数为1

```shell
# go run goroutine1.go
Max: 4
waiting to finish
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
num finish
a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z
lower finish
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
upper finish
Finished
# 测试使用，以及不设置时，对比运行差别。
```

doc: http://localhost:6060/pkg/runtime/#GOMAXPROCS

case2：设置runtime的cpu核数为2

```shell
# go run goroutine1.go
Max: 4
waiting to finish
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 A B C D E 2 3 4 5 6 7 8 9 F G H I J a b c d e f g h i j k l m n o p q r s t u v w x y
num finish
z a b c d e f g h i j k l m n o K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y p q r s t u v w x y z
lower finish
Z
upper finish
Finished
```



关于go一个函数时，函数的定义。可以是直接go一个匿名函数。也可以先定义好函数。再go该函数。下面的例子，更适合直接传参数。而goroutine1.go中可以直接wg而不传参数，以闭包的方式使用。

goroutine2.go是以函数方式调用的。

```go
var i int
for i:=0;i<10;i++{
  go func(var x int) {
    // process i
  }(i)
}
```



### prime

计算质数：

```shell
# go run prime.go

```



计算较多的质数，观察单核与双核的调度：

只有一个核，但因为要设计的质数比较多，一个routine计算后会被换出。由另一个进行。

```shell
#### GOMAXPROCS(1)
# go run prime2.go
go B: 2
go B: 3
go B: 5
go B: 7
# ...

go B: 3181
go B: 3187
go A: 2
go A: 3

go A: 4493
go A: 4507
go A: 4513
go B: 3191
go B: 3203
# ...

go B: 4993
go B: 4999
B finished
go A: 4517
go A: 4519
# ...

go A: 4999
A finished
```

如果将GOMAXPROCS置为2，则输出结果是两个routine同时输出。

通过对比，goroutine的调度是在逻辑核上进行的。

单个核时：所有的routine在队列中等待。正在运行的Gr到阻塞时，会让出cpu，让队列中的Gr运行；而阻塞的Gr如果恢复出来，再进入队列。

多个核时：是单核时的并行处理。



## goto/break/contine

### goto的坑

* 死循环

```go
func main() {
LABEL:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		goto LABEL
	}
}
// 持续不断的输出0
```



* goto到label之间有变量声明

goto到后面的区间内，不要声明变量

测试

```shell
# go run label.go
# command-line-arguments
./label.go:10:11: goto End jumps over declaration of out at ./label.go:15:6
```

应该将out定义在for前面。否则，没有被声明。



### 区别

- break跳出某个标签
- contine是继续某个标签，变量会继续迭代
- goto是跳到某个标签，变量重新开始运行迭代



同样是计算prime的函数。使用continue与break的差别。

goto正常使用的示例: 

```shell
# diff continuePrime.go breakPrime.go
16c16
< 				continue next
---
> 				break next

# go run continuePrime.go
go B: 2
	 continue:out=3,in=2
go B: 3
	 continue:out=4,in=2
	 continue:out=5,in=2
	 continue:out=5,in=3
	 continue:out=5,in=4
go B: 5
	 continue:out=6,in=2
	 continue:out=7,in=2
### ...
```

goto错误使用的示例 breakPrime.go

```shell
# go run breakPrime.go
go B: 2
	 continue:out=3,in=2
go B: 3
	 continue:out=4,in=2
B finished
go A: 2
	 continue:out=3,in=2
go A: 3
	 continue:out=4,in=2
A finished
```



