## net/http

refer:

* https://golang.org/pkg/net/http/
* https://gowalker.org/net/http



### overview



```go
// Get, Head, Post, PostForm make HTTP requests.

resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

the client must close the response body when finished with it.

```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}

// defer for close
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...

```

http.Client for control over HTTP client headers, redirect policy.

http.Transport for control over proxies, TLS configuration, keep-alive, compression.

**ListenAndServe**

ListenAndServe starts an HTTP server with a given address and a handler. the handler is usually nil, which means to use **DefaultServeMux**.

Handle and HandleFunc add handles to DefaultServeMux.

```go
// add handler
http.Handle("/foo", fooHandler)

// add func
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

more control over the server's behavior is available by creating a custom Server:

```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```



## Constants

* common HTTP methods like GET, POST… in RFC 7231
* HTTP status code 

```go
const DefaultMaxHeaderBytes = 1 << 20 // 1 MB
// the maximum permitted size of the headers in an HTTP request. 
// This can be overridden by setting Server.MaxHeaderBytes.
```



## Variables



错误相关返回值变量

使用了&ProtocolError结构，该结构只有一个string变量：ErrorString，并有一个方法Error()。

```go
ErrShortBody = &ProtocolError{"entity body too short"}

// HTTP request parsing errors.
type ProtocolError struct {
	ErrorString string
}

func (err *ProtocolError) Error() string { return err.ErrorString }
// from: https://github.com/golang/go/blob/master/src/net/http/request.go
```

ResponseWriter.Write 返回值

```go
ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")
ErrHijacked = errors.New("http: connection has been hijacked")
ErrContentLength = errors.New("http: wrote more than the declared Content-Length")

// from errors package https://golang.org/src/errors/errors.go
func New(text string) error
// New returns an error that formats as the given text.
```

default

```go
var DefaultClient = &Client{}
// is the default Client and is used by Get, Head, and Post.

var DefaultServeMux = &defaultServeMux
// is the default ServeMux used by Serve.
```

defaultServeMux在那里定义的？



## func

```go
func CanonicalHeaderKey(s string) string
// returns the canonical format of the header key s. 
// For example, the canonical key for "accept-encoding" is "Accept-Encoding".

func DetectContentType(data []byte) string
// DetectContentType implements the algorithm described at 
// http://mimesniff.spec.whatwg.org/ to determine the Content-Type of the given data. 
// It considers at most the first 512 bytes of data. 
// DetectContentType always returns a valid MIME type

```

Error

```go
func Error(w ResponseWriter, error string, code int)
// replies to the request with the specified error message and HTTP code. 
// It does not otherwise end the request; 
// The error message should be plain text.
```

**handle**

```go
// The documentation for ServeMux explains how patterns are matched.

func Handle(pattern string, handler Handler)
// registers the handler for the given pattern in the DefaultServeMux. 

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// registers the handler function for the given pattern in the DefaultServeMux. 

func ListenAndServe(addr string, handler Handler) error
// 1) listens on the TCP network address addr and then 
// 2) calls Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives. 
// Handler is typically nil, in which case the DefaultServeMux is used.

func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
// ignore
```

example of server: 

* net-http_server.go



```go
func MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser
// is intended for limiting the size of incoming request bodies.
// prevents clients from accidentally or maliciously sending a large request and wasting server resources.

func NotFound(w ResponseWriter, r *Request)
// NotFound replies to the request with an HTTP 404 not found error.

func ParseHTTPVersion(vers string) (major, minor int, ok bool)
// parses a HTTP version string. "HTTP/1.0" returns (1, 0, true).

func ParseTime(text string) (t time.Time, err error)
// parses a time header (such as the Date: header), 
// trying each of the three formats allowed by HTTP/1.1: 
// TimeFormat, time.RFC850, and time.ANSIC.

```



about serve

```go
func Serve(l net.Listener, handler Handler) error
// 1 Serve accepts incoming HTTP connections on the listener l, 
// 2 creating a new service goroutine for each. 
// 3 The service goroutines read requests and then call handler to reply to them. 
// 4 Handler is typically nil, in which case the DefaultServeMux is used.

func ServeContent(w ResponseWriter, req *Request, name string, 
                  modtime time.Time, content io.ReadSeeker)
// ServeContent replies to the request using the content in the provided ReadSeeker. 
// The main benefit of ServeContent over io.Copy is that :
// it handles Range requests properly, sets the MIME type, 
// and handles If-Modified-Since requests.

func ServeFile(w ResponseWriter, r *Request, name string)
// replies to the request with the contents of the named file or directory.
```



### Client struct

struct

```go
type Client struct {
  Transport RoundTripper	
  // specifies the mechanism by which individual HTTP request are made.
  // if nil, DefaultTransport is used.
  
  CheckRedirect func(req *Request, via []*Request) error
  // specifies the policy for handling redirects.
  	// If CheckRedirect is nil, the Client uses its default policy,
    // which is to stop after 10 consecutive requests.
  
  Jar CookieJar
  // specifies the cookie jar.
   	// If Jar is nil, cookies are not sent in requests and ignored in responses.
  
  Timeout time.Duration
  // specifies a time limit for requests made by this Client. 
  // the timeout includes connection time, any redirects, and 
  // reading the response body.   
  // The timer remains running after Get, Head, Post, or Do return 
  // and will interrupt reading of the Response.Body.
  // 
  // A Timeout of zero means no timeout.
```

调用的接口。

```go
type RoundTripper interface {
   RoundTrip(*Request) (*Response, error)
   // RoundTrip executes a single HTTP transaction, 
   // returning a Response for the provided Request.
}
```

a Client is an HTPP client.  Its **zero value (DefaultClient)** is a usable client that uses DefaultTransport.

The Client's Transport typically has internal state (cached TCP connections), so Clients should be **reused** instead of created as needed. Clients are safe for concurrent use by multiple goroutines.

### method of Client

```go
func (c *Client) Do(req *Request) (*Response, error)
// Do sends an HTTP request and returns an HTTP response, following policy (such as redirects, cookies, auth) as configured on the client.

func (c *Client) Get(url string) (resp *Response, err error)
// Get issues a GET to the specified URL. If the response is one of the  3xx redirect codes, Get follows the redirect after calling the Client's CheckRedirect function.
// 
// To make a request with custom headers, use NewRequest and Client.Do.

func (c *Client) Head(url string) (resp *Response, err error)
// issues a HEAD to the specified URL.  following redirect codes,

func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
// issues a POST to the specified URL.
// Caller should close resp.Body when done reading from it.
// To set custom headers, use NewRequest and Client.Do.

func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
// issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.
```

### CloserNotifier interface

```go
type CloseNotifier interface {
    // CloseNotify returns a channel that receives at most a
    // single value (true) when the client connection has gone away.
    //
    // CloseNotify may wait to notify until Request.Body has been
    // fully read.
    //
    // After the Handler has returned, there is no guarantee
    // that the channel receives a value.
    CloseNotify() <-chan bool
}
```

The CloseNotifier interface is implemented by **ResponseWriters** which allow detecting when the underlying connection has gone away.

This mechanism can be used to cancel long operations on the server if the client has disconnected before the response is ready.

### type connState

```go
type ConnState int
// A ConnState represents the state of a client connection 
// to a server. It's used by the optional Server.ConnState hook.

const (
	StateNew ConnState = iota
  	// represents a new connection that is expected to send a request immediately.
	// transition to either StateActive or StateClosed.

  	StateActive
  	// represents a connection that has read 1 or more bytes of a request. 
  	// The Server.ConnState hook for StateActive fires 
  	// before the request has entered a handler
  	// After the request is handled, the state transitions to 
  	// StateClosed, StateHijacked, or StateIdle.

	StateIdle
  	// represents a connection that has finished handling a request and 
  	// is in the keep-alive state, waiting for a new request. 
  	// Connections transition from StateIdle to either StateActive or StateClosed.

  	StateHijacked
	// represents a hijacked connection.
    // This is a terminal state. It does not transition to StateClosed.
        
	StateClosed
    // represents a closed connection. This is a terminal state. 
  	// Hijacked connections do not transition to StateClosed.
)
```

method of connState

```go
func (c ConnState) String() string
```



### Cookie 

```go
type Cookie struct {
   Naming, Value string
   MaxAge int
   Secure, HttpOnly bool
   Raw string
   // ...
}
// A Cookie represents an HTTP cookie as sent in the Set-Cookie header
// of an HTTP response or the Cookie header of an HTTP request.

func (c *Cookie) String() string

type CookieJar interface {
	//handles the receipt of the cookies in a reply for the given URL.
  	SetCookies(u *url.URL, cookies []*Cookie)
  
  	// returns the cookies to send in a request for the given URL.
    Cookies(u *url.URL) []*Cookie
}
// A CookieJar manages storage and use of cookies in HTTP requests.

```



### Dir, File, FileSystem



```go
type Dir string
// A Dir implements FileSystem using the native file system restricted to a specific directory tree.

func (d Dir) Open(name string) (File, error)

type File interface {
    io.Closer
    io.Reader
    io.Seeker
    Readdir(count int) ([]os.FileInfo, error)
    Stat() (os.FileInfo, error)
}
// A File is returned by a FileSystem's Open method and 
// can be served by the FileServer implementation.

// The methods should behave the same as those on an *os.File.

type FileSystem interface {
    Open(name string) (File, error)
}
// A FileSystem implements access to a collection of named files. 
// The elements in a file path are separated by slash ('/', U+002F) characters, 
// regardless of host operating system convention.

type Flusher interface {
    // Flush sends any buffered data to the client.
    Flush()
}
// The Flusher interface is implemented by ResponseWriters that 
// allow an HTTP handler to flush buffered data to the client.
```

Note that even for ResponseWriters that support Flush, if the client is connected through an HTTP proxy, the buffered data may not reach the client until the response completes.



## Handler

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

A Handler **responds** to an HTTP request.

ServeHTTP should write **reply headers and data** to the ResponseWriter and then **return**. 

Returning signals that the request is finished; it is not valid to use the ResponseWriter or read from the Request.Body after or concurrently with the completion of the ServeHTTP call. 

译：当对ServeHTTP的调用结束时，再使用ResponseWriter或者从Request.Body读取内容则是非法的。

it may not be possible to read from the **Request.Body** after writing to the ResponseWriter. Cautious handlers should read the Request.Body first, and then reply.

Except for reading the body, handlers should not modify the provided Request.

If ServeHTTP panics, the server (the caller of ServeHTTP) assumes that the effect of the panic was isolated to the active request. It recovers the panic, logs a stack trace to the server error log, and **hangs up**(挂断，搁置) the connection.

**FileServer**

```go
func FileServer(root FileSystem) Handler
// returns a handler that serves HTTP requests with 
// the contents of the file system rooted at root.

// To use the operating system's file system implementation, use http.Dir:
http.Handle("/", http.FileServer(http.Dir("/tmp")))

```

example :

* net-http_fileserver.go

```shell
# server
go run net-http_fileserver.go

# client
curl localhost:8080
<pre>
<a href="bash/">bash/</a>
<a href="cups/">cups/</a>
<a href="groff/">groff/</a>
<a href="ntp/">ntp/</a>
<a href="postfix/">postfix/</a>
</pre>
```



other func about handler

```go
func NotFoundHandler() Handler
func RedirectHandler(url string, code int) Handler

func StripPrefix(prefix string, h Handler) Handler
// StripPrefix returns a handler that serves HTTP requests by removing the 
// given prefix from the request URL's Path and invoking the handler h. 
```

example:

* net-http_stripFileserver.go

```shell
# server 
go run net-http_stripFileserver.go
# route /home/* to /Users/user/*

# client
curl localhost:8080/
# 404 page not found

curl localhost:8080/home
# <a href="/home/">Moved Permanently</a>.

curl localhost:8080/home/
<pre>
<a href=".Trash/">.Trash/</a>
<a href=".bash_history">.bash_history</a>
...
<a href=".vimrc">.vimrc</a>
<a href="Applications/">Applications/</a>
<a href="Desktop/">Desktop/</a>
<a href="Documents/">Documents/</a>
<a href="Downloads/">Downloads/</a>
</pre>

curl localhost:8080/home/
<pre>

curl localhost:8080/home/_go/src
# return blank

curl localhost:8080/home/_go/src/
<pre>
<a href="github.com/">github.com/</a>
<a href="golang.org/">golang.org/</a>
<a href="gopl.io/">gopl.io/</a>
<a href="readme.md">readme.md</a>
</pre>
```



**HandlerFunc, ServeHTTP**

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)
// 将一个函数f转换为一个Handler，并通过Handler来调用f。

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
// ServeHTTP calls f(w, r).

// HandlerFunc是一次函数转换，ServeHTTP是前者的一个方法。
```



### Header

```go
type Header map[string][]string
// A Header represents the key-value pairs in an HTTP header.

func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
// Get gets the first value associated with the given key. 
// If there are no values associated with the key, Get returns "". 
// To access multiple values of a key, 
// access the map directly with CanonicalHeaderKey.

func (h Header) Set(key, value string)
// Set sets the header entries associated with key to the single element value. It replaces any existing values associated with key.

func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error

```



### Hijacker

```go
type Hijacker interface {
  // Hijack lets the caller take over the connection.
  // After a call to Hijack(), the HTTP server library
  // will not do anything else with the connection.
  // It becomes the caller's responsibility to manage
  // and close the connection.
  Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```

The Hijacker interface is implemented by **ResponseWriters** that allow an HTTP handler to take over the connection.

example: 

* net-http_hijack.go

```shell
# server
go run net-http_hijack.go

# client
curl localhost:8080/hijack

# test1: 1) client input xyz; 2) server CTRL-C, 3) see client output
# test2: 1) client CTRL-C, 2) see server output
```



## type Request

struct:

A Request represents an HTTP **request** 

* received by a server or 
* to be sent by a client.



```go
type Request struct {
    Method string

    URL *url.URL

    // The protocol version for incoming server requests.
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0

    Header Header

    Body io.ReadCloser

    ContentLength int64

    TransferEncoding []string
    Close bool
    Host string
    Form url.Values
    PostForm url.Values
    MultipartForm *multipart.Form

    Trailer Header

    RemoteAddr string

    RequestURI string
    TLS *tls.ConnectionState

    Cancel <-chan struct{}

    // Response is the redirect response which caused this request
    // to be created. This field is only populated during client
    // redirects.
    Response *Response
    // contains filtered or unexported fields
}
```



### func about Request

```go
func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
// returns a new Request given a method, URL, and optional body.

func ReadRequest(b *bufio.Reader) (*Request, error)
// ReadRequest reads and parses an incoming request from b.

```

### Method about Request

```go
func (r *Request) AddCookie(c *Cookie)
func (r *Request) Cookie(name string) (*Cookie, error)
func (r *Request) Cookies() []*Cookie
//  parses and returns the HTTP cookies sent with the request.

func (r *Request) BasicAuth() (username, password string, ok bool)
// returns the username and password provided in the request's Authorization header

func (r *Request) Context() context.Context

func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
func (r *Request) FormValue(key string) string

// ...

func (r *Request) Write(w io.Writer) error

```



## Response

```go
type Response struct {
    Status     string // e.g. "200 OK"
    StatusCode int    // e.g. 200
    Proto      string // e.g. "HTTP/1.0"
    ProtoMajor int    // e.g. 1
    ProtoMinor int    // e.g. 0

  	Header Header
  	Body io.ReadCloser
  	ContentLength int64
  	TransferEncoding []string
  	//...
}
```

### funcs

```go
func Get(url string) (resp *Response, err error)
// issues a GET to the specified URL.
// Get is a wrapper around DefaultClient.Get.

func Head(url string) (resp *Response, err error)

func Post(url string, bodyType string, body io.Reader) (resp *Response, err error)

func PostForm(url string, data url.Values) (resp *Response, err error)

func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)

```

### methods

```go
func (r *Response) Cookies() []*Cookie

func (r *Response) Location() (*url.URL, error)
// returns the URL of the response's "Location" header

func (r *Response) Write(w io.Writer) error
// writes r to w in the HTTP/1.x server response format
```



## ResponseWriter interface

```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(int)
```



## ServeMux

```go
type ServeMux struct {
        // contains filtered or unexported fields
}

// func
func NewServeMux() *ServeMux
```

ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

Longer patterns take precedence over shorter ones

```go
// method
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
// returns the handler to use for the given request

func (mux *ServeMux) Handle(pattern string, handler Handler)
// registers the handler for the given pattern. 
// If a handler already exists for pattern, Handle panics.

func (mux *ServeMux) HandleFunc(pattern string, 
      handler func(ResponseWriter, *Request))
// registers the handler function for the given pattern.

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
// dispatches the request to the handler whose pattern 
// most closely matches the request URL.
```



## type Server

```go
type Server struct {
  Addr         string        
  // TCP address to listen on, ":http" if empty
  Handler      Handler       
  // handler to invoke, http.DefaultServeMux if nil
  ReadTimeout  time.Duration 
  // maximum duration before timing out read of the request
  WriteTimeout time.Duration 
  // maximum duration before timing out write of the response
  TLSConfig    *tls.Config   
  // optional TLS config, used by ListenAndServeTLS

  MaxHeaderBytes int
  TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
  ConnState func(net.Conn, ConnState)
  ErrorLog *log.Logger
}
// A Server defines parameters for running an HTTP server.

```

methods:

```go
func (srv *Server) ListenAndServe() error
// listens on the TCP network address srv.Addr and then 
// calls Serve to handle requests on incoming connections.

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error

func (srv *Server) Serve(l net.Listener) error
// accepts incoming connections on the Listener l, 
// creating a new service goroutine for each. 
// The service goroutines read requests and then 
// call srv.Handler to reply to them.

func (srv *Server) SetKeepAlivesEnabled(v bool)

```



## type Transport

Transport is an implementation of RoundTripper that supports HTTP, HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).

```go
// ingore
```

