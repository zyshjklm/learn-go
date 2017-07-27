
### 互斥锁

```shell

# 两个方法，Buy(), GiveWife()同时修改一个变量。
# 都能立即修改。
go run account-base.go
-1

# 设置bool变量flag标志，为真是可以修改，并立即关闭其他人可写，直到修改完成。
go run account-flag.go
5

# flag变成sync.Mutex. 使用Lock(), Unlock()
go run account-sync.go
5

# 由channel来代替main中的time.Sleep
# 通过协程封装消费方法，以及channel
# 这个结果应该不完全确定。
go run account-sync-chan.go
5

```

