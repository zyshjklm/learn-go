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



#### 5 Walk()

go error:

```go
for i := range ch {
    fmt.Println(i)
}
// fatal error: all goroutines are asleep - deadlock!
// beause: there is no close(ch) to close channel ch.

var c1 int
c1 <- ch1
// invalid operation: c1 <- ch1 (send to non-chan type int)
// because: you should write an assignment: like:
c1 = <- ch1

```



## end

** Where to Go from here...**

The [Go Documentation](http://golang.org/doc/) is a great place to start. It contains references, tutorials, videos, and more.

To learn how to organize and work with Go code, watch [this screencast](http://www.youtube.com/watch?v=XCsL89YtqCs) or read [How to Write Go Code](http://golang.org/doc/code.html).

If you need help with the standard library, see the [package reference](http://golang.org/pkg/). For help with the language itself, you might be surprised to find the [Language Spec](http://golang.org/ref/spec) is quite readable.

To further explore Go's concurrency model, watch [Go Concurrency Patterns](http://www.youtube.com/watch?v=f6kdp27TYZs) ([slides](http://talks.golang.org/2012/concurrency.slide)) and [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw) ([slides](http://talks.golang.org/2013/advconc.slide)) and read the [Share Memory by Communicating](http://golang.org/doc/codewalk/sharemem/)codewalk.

To get started writing web applications, watch [A simple programming environment](http://vimeo.com/53221558) ([slides](http://talks.golang.org/2012/simple.slide)) and read the [Writing Web Applications](http://golang.org/doc/articles/wiki/) tutorial.

The [First Class Functions in Go](http://golang.org/doc/codewalk/functions/) codewalk gives an interesting perspective on Go's function types.

The [Go Blog](http://blog.golang.org/) has a large archive of informative Go articles.

Visit [golang.org](http://golang.org/) for more.

