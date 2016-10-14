
## 1 about http.ListenAndServe("80", nil)的调用过程

after bind a handler for URI. use http.ListenAndServe() to start server.

```go
// github.com/golang/go/blob/master/src/net/http/server.go

// ListenAndServe is a func: 
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

// ListenAndServe is a method of Server
func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" { addr = ":http"}
	ln, err := net.Listen("tcp", addr)
	if err != nil { return err}
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```

### 1.1 net.Listen()

```go
// from net/dial.go

// Listen announces on the local network address laddr.
// The network net must be a stream-oriented network: "tcp", "tcp4",
func Listen(net, laddr string) (Listener, error) {
  	var l Listener
  	switch la := addrs.first(isIPv4).(type) {
	case *TCPAddr:
		l, err = ListenTCP(net, la)
	case *UnixAddr:
		l, err = ListenUnix(net, la)
	}
	return l, nil
}
```

Server的ListenAndServe函数，先调用net.Listen来监控一个本地端口，返回的是一个Listener接口。该接口有三个方法：

```go
// net/net.go

type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)
    // Close closes the listener.
    Close() error
    // Addr returns the listener's network address.
    Addr() Addr
}
```

Listen()又调用了ListenTCP(net, la)

```go
// from net/tcpsock.go

func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)

type TCPListener struct {
    fd *netFD
}	// TCPListener is a TCP network listener. 

// from net/fd_unix.go
type netFD struct {
   // ignore here...
}
```

其中返回值net.TCPListener是一个net包中定义的一个结构。而其中的指针，又是net/fd_unix.go中定义的一个描述符结构。


### 1.2 tcpKeepAliveListener

在收到来自net.Listen的net.TCPListener后，进入tcpKeepAliveListener函数。

```go
// from func (srv *Server) ListenAndServe() error {

return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
```

会话保持 tcpKeepAliveListener也是一个结构。且某成员也是一个TCPListener指针。

```go
// net/http/server.go
type tcpKeepAliveListener struct {
	*net.TCPListener
}
```


**tcpKeepAliveListener{  ln.(*net.TCPListener)  } —— 这是个什么表达式？**

Listener接口变量ln转化为tcpKeepAliveListener类型的监听，赋值给tcpKeepAliveListener匿名变量，最终调用了Server结构的Serve方法。

