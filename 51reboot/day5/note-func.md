## 过程编程



基础信息

basse/*



递归调用

recurse/*



函数返回值

type/*



匿名函数

anonimous/*



Return code

```shell

go run return.go
open abcd.txt: no such file or directory
exit status 2

go build return.go

./return
open abcd.txt: no such file or directory

echo $?
2
```



## 闭包

函数及其上下文。

一个函数被返回后，带有上下文。该函数可以使用在其函数以外定义的相关环境和参数。

闭包的坑：

code:closure-hole.go

第一轮使用循环时，i的地址都是一个，变的只是i的值。故最终结果是相同的。

第二轮使用循环时，通过i := i //给i变量重新赋值，每次循环时使用的是不同的地址。





错误处理

error/err*.go

error/panic*.go



## 命令执行



sys-call/exec-cmd.go

```go
out, err := cmd.CombinedOutput()
//在系统调执行该命令，把标准输出和错误输出都交给out。当然此时的它还是bytes类型.
```



sys-call/shell2.go

```go
fmt.Sscan(line, &cmd, &params)
// 通过Sscan获取到命令与参数。这时参数可能是空字符串。
// 将空串传给Command函数。会导致读不到执行结果。因此用了if len(params) == 0来判断
```

sys-call/shell3.go

```go
args := strings.Fields(line)
// 解决上述说的空串问题。
cmd := exec.Command(args[0], args[1:]...)	
// 因为当有空串时，Command的第二个参数是不存在的。为nil。
```



## 大整数

fib-iter/main.go

通过大整数及闭包来快速计算fibnacci数列。这里是顺序计算多个值，后者依赖于前者的计算结果。

fib-bigint/main.go

将前面的计算结果存储到map中，方便后继的计算直接读取已有结果，而不是从新计算一次。



