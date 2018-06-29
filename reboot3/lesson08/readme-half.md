2018 0623 下午

## 原子操作

atom1.go

核心代码

```go
for i := 0; i < 2; i++ {
    value := counter
    runtime.Gosched()
    value++
    counter = value
}
```

关于Gosched()

```powershell
# Gosched yields(放弃) the processor, allowing other goroutines to run. 
# It does not suspend the current goroutine, 
# so execution resumes automatically.
```

当routine中遇到这个函数时，routine将让出cpu时间片。

```shell
# go build atom1.go

# ./atom1
hello
out: 3

# ./atom1
hello
out: 2
```

输出的结果，时而为2，时而为3；有人的结果一直为4。

在main上增加`   runtime.GOMAXPROCS(1)`，限制可用计算资源。结果2，3，4都有。



因此，需要确保对counter的操作是原子的。

参考：http://localhost:6060/pkg/sync/atomic/

atom2.go

```shell
# go run atom2.go
hello
out: 4
```

atomic只是针对数的原子操作。如果是其他更复杂的多行代码的操作。需要其他机制。



atom3.go

测试`runtime.GOMAXPROCS(2)`改成不同的值，将for改成死循环。观察运行时的cpu使用情况。特别是不同的核。

通常设置为某个数，则相应数量的cpu核的cpu利用率将比较高，甚至达到100%。

