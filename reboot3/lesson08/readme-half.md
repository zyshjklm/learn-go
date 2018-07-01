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



## channel

channel是线程安全的。分为带缓冲和不带缓冲。

routine通过管道来共享资源。

### 无缓冲区

#### testChan

通过cobra进行封装。

```shell
# mkdir testChan
# cd testChan
# cobra init
Your Cobra application is ready at ... reboot3/lesson08/testChan

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.

# cobra add ball

# go run main.go ball
ball called
```

ball用于打球ball操作:

* 修改Run变量的内容
* init()中增加随机数种子
* 实现player函数

效果：

```shell
# go run main.go ball
ball called
start playing!!
chen starting!
Player chen hit ball 1 with rand 14
song starting!
Player song hit ball 2 with rand 60
Player chen hit ball 3 with rand 51
Player song hit ball 4 with rand 24
Player chen hit ball 5 with rand 13
Player song hit ball 6 with rand 48
Player chen hit ball 7 with rand 34
song miss, the number is 38
chen won!!!
```



无缓冲方式实现runner

```shell
# cobra add runner

# vim cmd/runner.go

# go run main.go runner
runner called
start runner()

runner 1 running with Baton
runner 1 use 0 seconds to the line
runner 1 exchange with runner 2

runner 2 running with Baton
runner 2 use 3 seconds to the line
runner 2 exchange with runner 3

runner 3 running with Baton
runner 3 use 3 seconds to the line
runner 3 exchange with runner 4

runner 4 running with Baton
runner 4 use 2 seconds to the line
runner 4 finish. Race over!

```



#### 封装unbuf包

* ball

```shell
# cd ../ 
# mkdir unbuf && cd unbuf
# vim player.go
#### 将上述ball相关的功能实现在player.go
# vim player_test.go

# go test
start playing!!
song starting!
song miss, the number is 38
chen starting!
chen won!!!
PASS
ok  	github.com/jkak/learn-go/reboot3/lesson08/unbuf	0.007s

```

* runner

```shell
# vim runner.go
# vim runner_test.go

# go test

```



#### testChan2调用unbuf

初始化环境

```shell
# cd ../
# mdkir testChan2 && cd testChan2

# cobra init
# cobra add ball
# cobra add runner

# go run main.go
```



修改cmd/ball.go, cmd/runner.go，调用unbuf包。

```shell
# go run main.go
# go run main.go runner
# go run main.go ball
```



### 有缓冲区

无缓冲区的channel，就像快递员到你家送信，你不在家就等着你，直到你亲自签收。

而有缓冲区的channel，就像你家门口的信箱。不管你在不在家（是否读取channel），只要信箱没有满，快递员直接将信放入信箱就走了。



### channel死锁

重点需要关注一下死锁。

实现一个功能，从一个chan读数据，然后将读到的数据写入另一个chan。

```shell
# go run main1.go
1
2
3
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
```

如果改变文件中for循环的次数，报警会略有不同。



分析这个main1.go文件，对channel的操作则2个关于channel读和写的代码：

```go
// write to ch2 
go func() {
    x := <-ch1
    ch2 <- x
}()

// read from ch2
for val := range ch2 {
    fmt.Println(val)
}
```

问题出现在for range会一直等待ch2，无法结束，这就会被Go判定产生了死锁。

