
### 互斥锁

```shell

# 两个方法，Buy(), GiveWife()同时修改一个变量。
# 都能立即修改。
go run account-base.go
-1

# 设置bool变量flag标志，为真是可以修改，并立即关闭其他人可写，直到修改完成。
go run account-flag.go
5

#### 使用flag标志变量，存在一些问题：
# 一个是这个变量本身是可能存在多个人同时修改
# 另外是等待变量切换状态，如果太频繁则耗资源，形成busy loop，太久则时间不浪费。

# flag变成sync.Mutex. 使用Lock(), Unlock()
# Mutex是互斥锁。有且只有一个协程能抢到锁。没抢到的则等待，直到锁释放。
go run account-sync.go
5

# 由channel来代替main中的time.Sleep
# 通过协程封装消费方法，以及channel
# 这个结果应该不完全确定。
go run account-sync-chan.go
5
```

wait

```shell

go run account-wait.go
5
go run account-wait.go
4
# 使用waitGroup来代替channel.
# 协程封装消费方法，以及waitGroup


# waitGroup与channel的差别在于后者可以设置超时。
# 而不用一直等等结束。

# 启动相应的协程，通过channel返回结果。
# 主协程使用select来选择channel的返回值,
# 以及定时器的到达，任意case满足则结束等待。
go run account-timeout.go
5
go run account-timeout.go
4
```

