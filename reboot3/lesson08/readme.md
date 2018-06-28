
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





