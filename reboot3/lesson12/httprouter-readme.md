

## httprouter



### 1) 基本使用

代码差异：

handle函数多一个router的参数

```go
// origin
func handler(w http.ResponseWriter, req *http.Request) {}

// router
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}
```

测试

```shell
# go run httprouter/main1.go

# curl localhost:7878/hello/
404 page not found

# curl localhost:7878/hello/a
hello, a!

# curl localhost:7878/
Welcome!
```



### 2) BasicAuth

对某些请求，请添加基础认证。如果不满足则登陆，满足之后才进入受保护的代码。

```go
router.GET("/protected/", BasicAuth(Protected, user, pass))

func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
  // core code: use 
  r.BasicAuth()
}

```

测试

```shell
# go run httprouter/main2.go

### 1）中的3项测试，这里结果一样。

# curl -i localhost:7878/protected/
HTTP/1.1 401 Unauthorized
Content-Type: text/plain; charset=utf-8
Www-Authenticate: Basic realm=Restricted
X-Content-Type-Options: nosniff
Date: Fri, 03 Aug 2018 13:40:44 GMT
Content-Length: 13

Unauthorized

# curl -i localhost:7878/protected/ -u jungle:secret
HTTP/1.1 200 OK
Date: Fri, 03 Aug 2018 13:41:01 GMT
Content-Length: 21
Content-Type: text/plain; charset=utf-8

Protected code here!
```

