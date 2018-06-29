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



## 互斥锁mutex

mutex1.go

通过使用sync.Mutex后，就不需要`runtime.GOMAXPROCS()`了。

Lock与Unlock之间的，就是临界区，保证同一时间只有一个routine可以操作临界区。

```shell
# go run mutex1.go
out: 4
```



### map读写

mutex2.go同时操作map的读和写。出现报错：

```powershell
# go run mutex2.go
fatal error: concurrent map iteration and map write

goroutine 5 [running]:
runtime.throw(0x10c5624, 0x26)
### ...
created by main.main
```

map不允许同时进行读写操作。map不是线程安全的。



mutex3.go

解决上述问题。将写操作设置为临界区，并加互斥锁。

但有时也会遇到fatal

```shell
# go run mutex3.go
fatal error: concurrent map iteration and map write

goroutine 5 [running]:
runtime.throw(0x10c5664, 0x26)

# go run mutex3.go
over

# go run mutex3.go
1 1
2 2
6 6
8 8
0 0
3 3
4 4
5 5
7 7
9 9
over
```



### 读写锁

mutex4.go

```shell
# go run mutex4.go
== write == 0
== write == 1
== write == 2
== write == 3
== write == 4
-- read id=4: key=0, val=0
-- read id=4: key=1, val=1
-- read id=4: key=2, val=2
### ...
-- read id=1: key=4, val=4
== write == 5
-- read id=18: key=0, val=0
-- read id=9: key=0, val=0
-- read id=4: key=4, val=4
-- read id=4: key=0, val=0
-- read id=4: key=1, val=1
-- read id=4: key=2, val=2
-- read id=4: key=3, val=3
-- read id=18: key=1, val=1
-- read id=17: key=3, val=3
-- read id=17: key=4, val=4
```

有写操作时，需要等所有的读完成，等其它的写完成。

有读操作时，多个读之间无影响。也就是上述输出中，id是乱序的。



mutex5.go

多个读协程可以同时读。如果用互斥锁，则多个读协程需要抢锁。

在输出上，读写锁的id是乱序的。互斥锁，则是同一个id抢到锁，同id的输出是一起的。

```shell
# go run mutex5.go

-- read id=17: key=1, val=1
-- read id=17: key=2, val=2
-- read id=17: key=3, val=3
-- read id=17: key=4, val=4
-- read id=17: key=0, val=0
-- read id=10: key=0, val=0
-- read id=10: key=1, val=1
-- read id=10: key=2, val=2
-- read id=10: key=3, val=3
-- read id=10: key=4, val=4
-- read id=11: key=0, val=0
-- read id=11: key=1, val=1
-- read id=11: key=2, val=2
-- read id=11: key=3, val=3
-- read id=11: key=4, val=4
```



