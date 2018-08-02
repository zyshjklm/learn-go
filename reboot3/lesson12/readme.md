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



#### mux多路选择器

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

