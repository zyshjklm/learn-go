



[refer from](https://segmentfault.com/a/1190000014875956)

### 拆分文件



```shell
curl localhost:8080/
Welcome!%

curl localhost:8080/todos/asd
Todo Show:asd%

curl localhost:8080/todos
[{"name":"Write presetation","completed":false,"due":"0001-01-01T00:00:00Z"},{"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"}]
```



### 添加日志

```shell
go run *go
2018/06/12 08:19:11 GET	/todos/12	TodoShow	10.184µs
2018/06/12 08:19:18 GET	/todos	TodoIndex	227.959µs

```



### 拆分route代码

见todos2/



### POST数据

增加repo.go模拟DB。

增加上传数据的路由与逻辑

```shell
# POST
curl -v -H "Content-Type: application/json" -d '{"name": "New Todo"}' http://localhost:8080/todos

> POST /todos HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.49.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 21
>
* upload completely sent off: 21 out of 21 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json; charset=UTF-8
< Date: Tue, 12 Jun 2018 14:37:47 GMT
< Content-Length: 75
<
{"id":3,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}

# get
curl -v  localhost:8080/todos
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /todos HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.49.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Tue, 12 Jun 2018 14:37:35 GMT
< Content-Length: 237
<
[{"id":1,"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":2,"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":3,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}]
```

