## concurrency

a goroutine is a lightweight thread managed by the Go runtime.

```go
go f(x, y, z)
// starts a new goroutine running f(x, y, z)
```

the evaluation of f, x, y and z happens in the current goroutine and the execution of f happens in the new goroutine.

goroutines run in the same address space, so access to shared memory must be synchronized.

#### 1 channel

channels are a typed conduit(导管) through which you can send and receive values with the channel operator, `<-`.

```go
ch := make(chan int)
ch <- v		// send v to channel ch.
v := <-ch	// receiv from ch, and assignn value to v.
```

the data flows in the direction of the arrow.

channels must be created before use.

by default, sends and receives block until the other side is ready. this allows goroutines to synchronize without explicit lock or condition variables.

