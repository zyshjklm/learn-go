

### mysql driver

其中有init会初始化注册一个"mysql".
使用的是database.sql.Register("mysql", &xxxx)


### _方式引用一个包

引用一个包而又不用这个包，就可以使用这种方式。
实际上是使用了这个包中的init函数。


### sqlite3 

```go
{
    import _ "github.com/mattn/go-sqlite3"

    db, err := sql.Open("sqlite3", "web.db")
    if err != nil {
        log.Fatal(err)
    }
    db.Ping()
}
```

### checkLogin()

```shell

# CheckLogin usage:
curl http://localhost:8090/checkLogin?user=admin&password=admin

```
### Add()

usage 
```shell

curl http://localhost:8090/add?name=adminni&password=admin&isadmin=0&note=hello,administrator
# null
# 200 

curl http://localhost:8090/add?name=admin&password=admin&isadmin=0&note=hello,administrator
# Error 1062: Duplicate entry 'admin' for key 'name'
# 500


```

### msql table

mysql

```mysql
CREATE TABLE `user` (
`id`  int(11) UNSIGNED NOT NULL AUTO_INCREMENT ,
`name`  char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' ,
`password`  char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' ,
`note`  varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' ,
`isadmin`  tinyint(1) NOT NULL ,
PRIMARY KEY (`id`),
UNIQUE INDEX `name` (`name`) USING BTREE
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=6
ROW_FORMAT=COMPACT
;

```

### 多条记录查询

单条循环会比较慢，可以使用prepare，加快处理速度。


### debug pprof

http://localhost:8090/debug/
http://localhost:8090/debug/pprof/
http://localhost:8090/debug/pprof/goroutine?debug=1

```shell

go tool pprof http://localhost:8090/debug/pprof/heap
Fetching profile from http://localhost:8090/debug/pprof/heap
Saved profile in /Users/song/pprof/pprof.localhost:8090.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
Entering interactive mode (type "help" for commands)
(pprof) top
1.16MB of 1.16MB total (  100%)
Showing top 10 nodes out of 12 (cum >= 1.16MB)
      flat  flat%   sum%        cum   cum%
    1.16MB   100%   100%     1.16MB   100%  runtime/pprof.writeGoroutineStacks
         0     0%   100%     1.16MB   100%  github.com/gorilla/handlers.(*loggingHandler).ServeHTTP
         0     0%   100%     1.16MB   100%  github.com/gorilla/handlers.loggingHandler.ServeHTTP
         0     0%   100%     1.16MB   100%  net/http.(*ServeMux).ServeHTTP
         0     0%   100%     1.16MB   100%  net/http.(*conn).serve
         0     0%   100%     1.16MB   100%  net/http.HandlerFunc.ServeHTTP
         0     0%   100%     1.16MB   100%  net/http.serverHandler.ServeHTTP
         0     0%   100%     1.16MB   100%  net/http/pprof.Index
         0     0%   100%     1.16MB   100%  net/http/pprof.handler.ServeHTTP
         0     0%   100%     1.16MB   100%  runtime.goexit
(pprof)


```

热点分析

启动如下命令后，处于30秒收集信息中。
可以多次在浏览器中访问某些页面。如counter，list

```shell

go tool pprof http://localhost:8090/debug/pprof/profile
Fetching profile from http://localhost:8090/debug/pprof/profile
Please wait... (30s)
Saved profile in /Users/song/pprof/pprof.localhost:8090.samples.cpu.002.pb.gz
Entering interactive mode (type "help" for commands)
(pprof) top
20ms of 20ms total (  100%)
Showing top 10 nodes out of 14 (cum >= 10ms)
      flat  flat%   sum%        cum   cum%
      10ms 50.00% 50.00%       10ms 50.00%  runtime.kevent
      10ms 50.00%   100%       10ms 50.00%  runtime.mach_semaphore_wait
         0     0%   100%       20ms   100%  runtime.findrunnable
         0     0%   100%       10ms 50.00%  runtime.goexit0
         0     0%   100%       20ms   100%  runtime.mcall
         0     0%   100%       10ms 50.00%  runtime.netpoll
         0     0%   100%       10ms 50.00%  runtime.notesleep
         0     0%   100%       10ms 50.00%  runtime.park_m
         0     0%   100%       20ms   100%  runtime.schedule
         0     0%   100%       10ms 50.00%  runtime.semasleep
(pprof) top10
20ms of 20ms total (  100%)
Showing top 10 nodes out of 14 (cum >= 10ms)
      flat  flat%   sum%        cum   cum%
      10ms 50.00% 50.00%       10ms 50.00%  runtime.kevent
      10ms 50.00%   100%       10ms 50.00%  runtime.mach_semaphore_wait
         0     0%   100%       20ms   100%  runtime.findrunnable
         0     0%   100%       10ms 50.00%  runtime.goexit0
         0     0%   100%       20ms   100%  runtime.mcall
         0     0%   100%       10ms 50.00%  runtime.netpoll
         0     0%   100%       10ms 50.00%  runtime.notesleep
         0     0%   100%       10ms 50.00%  runtime.park_m
         0     0%   100%       20ms   100%  runtime.schedule
         0     0%   100%       10ms 50.00%  runtime.semasleep
(pprof) top20
20ms of 20ms total (  100%)
      flat  flat%   sum%        cum   cum%
      10ms 50.00% 50.00%       10ms 50.00%  runtime.kevent
      10ms 50.00%   100%       10ms 50.00%  runtime.mach_semaphore_wait
         0     0%   100%       20ms   100%  runtime.findrunnable
         0     0%   100%       10ms 50.00%  runtime.goexit0
         0     0%   100%       20ms   100%  runtime.mcall
         0     0%   100%       10ms 50.00%  runtime.netpoll
         0     0%   100%       10ms 50.00%  runtime.notesleep
         0     0%   100%       10ms 50.00%  runtime.park_m
         0     0%   100%       20ms   100%  runtime.schedule
         0     0%   100%       10ms 50.00%  runtime.semasleep
         0     0%   100%       10ms 50.00%  runtime.semasleep.func1
         0     0%   100%       10ms 50.00%  runtime.semasleep1
         0     0%   100%       10ms 50.00%  runtime.stopm
         0     0%   100%       10ms 50.00%  runtime.systemstack
(pprof)

```

### 路由分发器

ServeMux
DefauleServeMux

```go
    // style 1
	log.Fatal(http.ListenAndServe(":8090", nil))

    // style 2
	h := handlers.LoggingHandler(os.Stderr, http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8090", h))

```

### Handler and HandlerFunc

http.HandlerFunc 是函数，使用http.HandleFunc挂载。
http.Handler是接口，使用http.Handle挂载

转换关系：
比如Login()是一个处理函数，如何转换成接口来挂载
http.HandlerFunc(Login) -> http.Handler


