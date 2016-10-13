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
// It considers at most the first 512 bytes of data. DetectContentType always returns a valid MIME type

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

