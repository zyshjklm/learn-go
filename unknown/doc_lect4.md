## 4 const and operator

 字符串与整型的转换。

strconv.Itoa(xx),  int to char

strconv.Atoi(xx),  char to int



### 常量

常量的值在编译的时候就已经确定了。

常量表达式内部的函数必须时内置函数，这样在编译时才能知道其运行结果。

在定义常量组时，如果不提供初始值，则表示使用上一行的表达式

使用相同的表达式不代表具有相同的值



### iota常量计数器

iota计数从0开始，每定义一个常量自动递增1(此时常量未必使用了iota)，

iota的计数是从const的第一个常量就开始计数。

每定义一个const，计算数0。

 

### 新类型与别名

```go
type newint int

// newint并不是int的别名，而是一种自定义类型。
//在类型转换时，需要显示转换。但：

// byte和uint8是别名
// rune和int32是别名

```

