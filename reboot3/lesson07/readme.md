
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


