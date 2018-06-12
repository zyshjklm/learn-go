

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

