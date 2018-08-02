## lesson12



### http

**http.HandleFunc()**

http1.go

只有如下内容：

```go
func main() {
	log.Fatal(http.ListenAndServe(":7878", nil))
}
```

运行效果

```shell
# go run http1.go

# curl localhost:7878/
404 page not found
```

运行时启动了服务端口。但不提供任何服务。



http2.go

增加了处理函数

```shell
# go run http2.go

# curl localhost:7878/
hello golang...%
```

这两个例子都是使用的http.HandleFunc()来绑定路由。该函数是一个如下类型的函数：

`func handleFuncName(w http.ResponseWriter, req *http.Request)`



http3.go

使用http.Server来实例化服务。并绑定如下某个实例的固定名称的方法做为处理函数。

`ServeHTTP(w http.ResponseWriter, r *http.Request)`

```shell
# go run http3.go

# curl localhost:7878/
hello golang from myhandle...%
```



#### mux多路利用器

httpMux1.go

多个handler。基于http3.go。使用**http.Handle**来为http.Server加载路由。

3个对象都实现了ServeHTTP()方法。

```shell
# go run httpMux1.go

# curl localhost:7878/
hello root%                                                                                                            

# curl localhost:7878/world
hello world%                                                                                                          

# curl localhost:7878/hello
hello from golang%
```



httpMux2.go

与前者不同，本次使用http.HandleFunc来加载路由函数。而前者是某个对象的特定方法。

```shell
# go run httpMux2.go

# curl localhost:7878/hello
from hello...%                                                                                                         

# curl localhost:7878/world
from world...%                                                                                                         

# curl localhost:7878/
hello golang...%
```



多路复用

http.HandleFunc(pattern, handlerF)
底层的muxEntry做来的存储路由信息。



```go
http.Handle()       注册处理器，默认ServeHTTP方法。即某个对象的方法。
http.HandleFunc()   注册处理器函数。
DefaultServeMux
// ServeMux的Handler方法，用于根据pattern来进行方法的路由选择。

ServeMux匹配规则
// 将httpMux2.go中 
// http.HandleFunc("/hello", helloHandler) 改为：
// http.HandleFunc("/hello/", helloHandler) 观察效果
```

运行效果：

```shell
# curl localhost:7878/hello
<a href="/hello/">Moved Permanently</a>.

# curl localhost:7878/hello/
from hello...%

```



#### 装饰器

基于`http.HandleFunc`的装饰函数打印日志

```go
func logHTTP(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		p := runtime.FuncForPC(reflect.ValueOf(h).Pointer())
		name := p.Name()
		fmt.Println("handler func called = ", name)
		h(w, req)
	}
}
```

效果

```shell
# go run httpDecorator1.go
handler func called =  main.helloHandler

# curl localhost:7878/hello/
hello golang...%                                                                                                       

# curl localhost:7878/hello
from hello...%
```

当请求localhost:7878/hello时，会命中调用打印日志。



基于`http.Handler`接口的装饰函数

```go
func logHandler(h http.Handler) http.Handler {
	temp := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(temp)
}
```

注意这里与上一个示例的区别。

* http.HandlerFunc是一个type of `func(ResponseWriter, *Request)`
* http.HandlerFunc类型有一个ServeHTTP方法。该方法调用f(w,r)
* http.HandlerFunc(f(w,r))时，实际是：Handler that calls f

执行效果

```shell
# go run httpDecorator2.go
Handler called - *main.myhello

# curl localhost:7878/hello/
hello root%                                                                                                            

# curl localhost:7878/hello
hello from golang%
```



