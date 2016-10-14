
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

### 1.3 Serve方法

如上所述，最终在拿到一个tcpKeepAliveListener类型的监控连接后，做为参数调用了Server结构的Serv方法。


```go
// from net/http/server.go

// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each. The service goroutines read requests 
// and then call srv.Handler to reply to them.

// ignore some code here...
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure

	baseCtx := context.Background() // base is always background, per Issue 16220
	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	ctx = context.WithValue(ctx, LocalAddrContextKey, l.Addr())
	for {
		rw, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				// modify tempDelay here...
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve(ctx)
	}
}
```

Serve函数调用了监控者接口的Accept，以及Server的newConn(c net.Conn).

#### 1.3.1 Accept()

Listener只是一个interface，而本处的类型是来自tcpKeepAliveListener{ln.(*net.TCPListener)}。这两个类型的Accept函数都调用了一个accept()函数。

* TCPListener.Accept()
  * return **TCPListener.accept()**
* tcpKeepAliveListener.Accept()  // net/http/server.go
  * tc = net.TCPListener.AcceptTCP()
    * c = **TCPListener.accept()**
    * return c
  * tc.SetKeepAlive(true)
  * tc.SetKeepAlivePeriod(3*time.Minute)
  * return tc   // net.Conn

基于unix的TCP的accept()定义在：

```go
// from net/tcpsock_posix.go

func (ln *TCPListener) accept() (*TCPConn, error) {
        fd, err := ln.fd.accept()
        if err != nil { return nil, err }
        return newTCPConn(fd), nil
}

// from net/tcpsock.go
type TCPListener struct {
        fd *netFD
}
```

TCPListener的成员fdnet/tcpsock.go中的netFD结构的指针。netFD结构在fd_unix.go中：

```go
// from net/fd_unix.go

func (fd *netFD) accept() (netfd *netFD, err error) {
    if err := fd.readLock(); err != nil {
            return nil, err
    }
    defer fd.readUnlock()

    var s int
    var rsa syscall.Sockaddr
    if err = fd.pd.prepareRead(); err != nil {
            return nil, err
    }
    for {
            s, rsa, err = accept(fd.sysfd)
    }
}
```

这里竟然又调用了一层accept()，从返回值上看，应该是调用了：

```go
// from syscall/net_nacl.go

func (f *netFile) accept() (fd int, sa Sockaddr, err error) {
        msg, err := f.listener.read(f.readDeadline())
        if err != nil {
                return -1, nil, err
        }
        newf, ok := msg.(*netFile)
        if !ok {
                // must be eof
                return -1, nil, EAGAIN
        }
        return newFD(newf), newf.raddr.copy(), nil
}

```

#### 1.3.2 newConn()

再回到Server的Serve方法中，最后几行：

```go
  c := srv.newConn(rw)
  c.setState(c.rwc, StateNew) // before Serve can return
  go c.serve(ctx)
```

Accept返回一个TCPConn指针。然后调用newConn()

```go
// from net/http/server.go

// Create new connection from rwc.
func (srv *Server) newConn(rwc net.Conn) *conn {
	c := &conn{
		server: srv,
		rwc:    rwc,
	}
	if debugServerConnections {
		c.rwc = newLoggingConn("server", c.rwc)
	}
	return c
}
```

建立连接后，返回一个conn结构类型的指针。这个结构在

```go
from net/net.go

type conn struct {
    fd *netFD
}
```

这还是个网络描述符。但在http中，又包括了net并重新定义了conn：

```go
// from net/http/server.go

type conn struct {
   server *Server
   rwc net.Conn
   //...
}

// Serve a new connection.
func (c *conn) serve(ctx context.Context) {
	//... 
  	for {
       	w = c.readRequest(ctx)
  		serverHandler{c.server}.ServeHTTP(w, w.req)
  		//...
    }
}
```

上面的serverHandler是同一个文件中定义的结构。c是一个连接指针，其方法readRequest读取连接信息到w变量。再用c初始化一个serverHandler类型的Server指针，并调用serverHander的ServeHTTP方法。



#### 1.4 serverHandler

```go
type serverHandler struct {
    srv *Server
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    handler := sh.srv.Handler
    if handler == nil {
      	handler = DefaultServeMux
    }
    handler.ServeHTTP(rw, req)
}
```



这个ServeHTTP方法实现了**Handler接口**。

```go
// from net/http/server.go
type Handler interface {
  	ServeHTTP(ResponseWriter, *Request)
}

// 下面是最关键的地方了

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
}
```



## 2 HandleFunc 

style 1: 

```go
// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    log.Fatal(http.ListenAndServe(":12345", nil))
}
```

如之前分析。ListenAndServe最终会调用到

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w, r)}
// 从而执行：
HelloServer(f, r)
```



## 3 Handle

style 2:

```go
package main

import (
    "log"
    "fmt"
    "net/http"
)

type String string

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s for URL.Path(%q)\n", str, r.URL.Path)
}

func main() {
    str := String("I'm jungle!")
    http.Handle("/string", str)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
```

example as read_pkg/web_server_handler.go

对于String类型，其实现了ServeHTTP，则实现了Handler接口。这样，在1.4部分所说的serverHandler在执行自己的ServeHTTP方法时，最终会找到默认的handler，并执行其自己定义的ServeHTTP方法。

```go
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    handler := sh.srv.Handler
    if handler == nil {
      	handler = DefaultServeMux
    }
    handler.ServeHTTP(rw, req)
}
```



因此：

使用HandleFunc时，显示的绑定处理函数xxxHandler(w, r)。然后通过HandlerFunc type的ServeHTTP来调用了xxxHandler(w, r)。

使用Handle来安装Handler时，需要显式实现自己的ServeHTTP函数，最终由serverHandler获取自定义Handler，因为自己实现的ServeHTTP实现了Handler接口，从而调用到自定义的ServeHTTP接口函数。



