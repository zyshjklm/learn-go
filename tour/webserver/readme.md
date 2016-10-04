## webserver

package http serves HTTP requests using any value that implements `http.Handler`:

```go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
```

相关函数原型：

```go
func ListenAndServe(addr string, handler Handler) error

func HandleFunc(pattern string, handler func(ResponseWriter, *Request) )

func Handle(    pattern string, handler Handler)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

#### 实现路由绑定：方式一

使用HandleFunc函数来绑定。将函数为一个请求的URI绑定一个处理函数，帮第二个参数是一个具体的函数。一般是先定义好这个函数。再将函数名放在这个HandleFunc函数中。

```go
func xxx_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

http.HandleFunc("/", xxx_handler)
log.Fatal(http.ListenAndServe("localhost:8000", nil))
```

这个方法比较简洁明了，路由对应处理函数就可以了。

这种方式背后，隐藏着一个东西，那就是Handler这个接口，Handler接口有个ServeHTTP方法。**HandleFunc接口自动包装实现了Handler的ServeHTTP方法**。

```go
// xxx_handler在被做为参数传给HandleFunc函数时，实际是这样子的
type HandleFunc xxx_handler(w, r) 

func (hf HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    hf.xxx_handler(w, r) 
}
```

所以另一个方法就是下面说的，显示的实现ServeHTTP方法。

#### 实现路由绑定：方式二

假设有一个struct或者某种自定义的类型，每次请求需要该类型的数据做为网页请求的响应值。则直接在该struct或类型上实现一个ServeHTTP方法，如此则**实现了Handler接口**。然后就可以直接使用http.Handle函数来绑定该struct或者类型的一个变量（因为这个变量所对应类型已经实现Handler接口）

```go
type String string

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "return %s for URL.Path(%q)\n", str, r.URL.Path)
}

str := String("I'm jungle!")
http.Handle("/string", str)
log.Fatal(http.ListenAndServe("localhost:8000", nil))

```

第一个方式是显示的绑定函数，但隐式的调用了ServeHTTP函数；传递的是某个函数；

第二个方式是隐式的绑定函数，但显式的实现了ServeHTTP函数，传递的是某类型的变量。

