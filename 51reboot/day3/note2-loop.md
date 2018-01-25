

### for 

* 基本循环式
* while式
* do while式
* 永久循环式

### swich

for/main.go

switch/main.go



### file open seek

使用了如下函数

* CreateFile
* Open
* Seek
* Read
* bufio.ReadString


示例文件：

* open/main.go
* read/main.go



### rand

rand/main.go

使用了：

* rand.Seed(time.Now().Unix())
* rand.Int63()
* strconv.FormatInt(x, 36)




#### fmt 乘法表

fmt/fmt.go



#### counter

实现2个操作数的加减乘除。

```shell
go run counter/main.go 10 \* 3
10 * 3
30

go run counter/main.go 10 / 0
10 / 0
fatal: div by zero!
```

counter/main.go



#### myls

实现一个简易的ls命令。

列出类型及大小。

myls/main.go

