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

#### 2 buffered channel

channels can be buffered. 

```go
ch := make(chan int, BUF_LEN)
```

sends to a buffered channel block only when the buffer is full.

receives block when the buffer is empty.

#### 3 range and close

a sender can `close` a channel to indicate that no more values will be sent. receivers can test whether a channel has been closed by assigning a second paramter to the receiver expression:

```go
v, ok := <-ch
```

ok is `false` if there are no more values to receive and the channel will closed. (sending on a closed channel will cause a panic.)

the loop `for i := range c` receives values from the channel repeatedly until it is closed.

closing channel only necessary when the receiver must be told there are no more values coming. such as to terminate a range loop.

#### 4 select

the `select` statement lefts a gorouting wait on multiple communication operations.

a `select` blocks until one of its case run, then it executes that case. it chooses one at random if multiple are ready.

