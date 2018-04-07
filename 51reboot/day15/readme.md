## 第15课

### 1 系统包 

atomic  原子变量

示例：atomic1.go

* 运行次数不一定。因为每个协程的修改不可预期。

示例：atomic2.go

* 使用锁。
* 示范使用原子操作。`atomic.AddInt32(&n, 2)`



```shell
### GOARCH 应用程序将要运行平台的处理器架构。影响指令集
### 	386, arm, amd64, ppc64, mips 等类型

### GOOS 应用程序将要运行平台的操作系统。
### 	linux, windows, darwin, freebsd, android等类型
# GOARCH=386 go build atomic1.go
# go build atomic1.go

# file atomic[12]
atomic1: Mach-O 64-bit executable x86_64
atomic2: Mach-O executable i386

```



### 2 unsafe 

unsafe是上帝视角。通过Pointer，可以代表任意类型。

```shell
# go run unsafe1.go
8

# GOARCH=386 go run unsafe1.go
4

# GOARCH=amd64 go run unsafe1.go
8

# GOARCH=mips go run unsafe1.go
cmd/go: unsupported GOOS/GOARCH pair darwin/mips

```

结构体的内存

```shell
# go run unsafe2.go
8
16
1
8
0
8
```

2个int8对齐成2字节；2个32对齐成8字节。但不同的类型会按更大的一个进行对齐。



#### 2.1 Pointer

一个int类型的指针，不能直接指向int64类型的变量；即使这两个类型在底层是一相大小。强制转换也不行。这样才能保证类型安全。

但在数字内容上是可以转换了。将一个类型的变量，转换成另一个类型并取值赋给新变量。

```shell
# go run unsafe3.go
32
68719476730
[0 0]
257
```



#### 2.2 slice

通过unsafe.Pointer来操纵切片。

```shell
# go run unsafe4.go
0xc4200181a0
main.SliceHeader{Data:(unsafe.Pointer)(0xc4200181a0), Len:3, Cap:3}
main.SliceHeader{Data:(unsafe.Pointer)(0xc4200181a0), Len:1, Cap:3}
```

人工构造slice

需要注意两个hdr赋值时的data字段的差别。导致的结果也不一样。使用Pointer需要小心。

* hdr2 使用的`&s[0]`做地址
* hdr3使用的`&s`做地址。导致切片内的第一个值不是1。

```shell
# go run unsafe5.go
0xc42001c060
3
6
1 2
end hdr2 test --:

3
6
842350575712 6
end hdr3 test --:

[1 2 3 4 5 6]
2
6
2 3
```

这里定义的切片数据是使用的`unsafe.Pointer`类型，golang中reflect中定义的SliceHeader中的data是`uintptr`类型。



#### 2.2 stringHeader

进行字符串的切片及拷贝。这个脚本有点难。

参考：unsafe6String.go。



### 3 反射reflect

用于动态的操纵数据。

示例：reflect1.go

新人容易使用refect，因为觉得很多东西都是不确定的。

go 建议事先设计好结构和字段。而不是让很多都不确定。当然像json的序列化，是通用性的，存在不确定性。当大多都确定时，就会少用reflect。因为reflect会影响性能。



操作json的marshal转换。没有处理最后多出来的`,`。

示例：reflect2.go



### 4 正则reg

简单使用。

代码：regexp.go



### 5 rpc

代码：

* rpc/client/main.go
* rpc/server/main.go

server:

```shell
go run client/main.go
2018/04/03 09:21:11 result:26
```

client

```shell
go run client/main.go
2018/04/03 09:21:11 result:26
```



需要注意的是，服务中的结构是个空对象。没有任何具体的字段。但有服务的方法Add()

```go
type MathService struct {
}
```

common中定义的两个结构，分别用来进行做为请求参数和返回结果。

client在调用时，通过`服务.方法`的方式来调用。即示例中的`MathService.Add`

rpc封装了底层的通信协议和传输过程。像调用`本地函数`一样调用网络服务。



### 6 protobuf

#### 6.1 编码

建议使用v3版本。在linux里，默认的可能是v2。所以考虑用源码进行安装。相关步骤：

* 安装protoc
  * 地址：https://github.com/google/protobuf/releases
  * 下载：[v3.5.0](https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-osx-x86_64.zip)
  * 解压：`unzip protoc-3.5.0-osx-x86_64.zip`
  * 得到：`{bin/protoc, include/}`
* 安装go的插件工具：
  * go get -v github.com/gogo/protobuf/proto
  * go get -v github.com/golang/protobuf/protoc-gen-go
* 定义proto
  * 编写**myproto/addressbook.proto**。这个是模板原型。
  * 类似于数据库中的表结构。类似于golang中的struct。
  * pb支持向下兼容，新增加的字段在老的接口中也能使用。
  * repeated相当于是golang里的slice。
* 生成pb

```shell
# bin/protoc --go_out=. myproto/addressbook.proto

# ls -trl myproto
-rw-r--r--  1 song  staff   347 Apr  6 23:36 addressbook.proto
-rw-r--r--  1 song  staff  5424 Apr  6 23:36 addressbook.pb.go

# vim encode.go

# go run encode.go > jungle.pb
```

#### 6.2 解码

```shell
# go run encode.go > jungle.pb

# go run decode.go < jungle.pb
{1 jungle jungle@github.com [number:"18612349876"  number:"87651234" type:HOME ]}
id:1 name:"jungle" email:"jungle@github.com" phones:<number:"18612349876" > phones:<number:"87651234" type:HOME >

```



#### 6.3 bench测试

比较proto与json的编码速度

```shell
# vim encode_test.go

# go test -bench .
# github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf
./encode.go:10:6: main redeclared in this block
	previous declaration at ./decode.go:12:6
FAIL	github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf [build failed]

##### 因为decode.go与encode.go在同一个目录下，且都是main。分开放到不同目录

# mkdir encode decode
# git mv encode.go encode
# mv encode_test.go encode

# mv decode.go decode

# cd encode
# go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/encode
BenchmarkProto-4   	 2000000	       705 ns/op
BenchmarkJSON-4    	 1000000	      1465 ns/op
PASS
ok  	github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/encode	3.640s

```

proto大概是json的2倍，如果数据再复杂些，量再多些，应该会差异更大。





### 7 grpc

#### 7.1 生成pb代码

```shell
go get -v google.golang.org/grpc

mkdir rpcrpto
vim rpcproto/addrbookstore.proto
#### 在proto/addrbook.proto 基础上增加了rpc相关message，记得改包名。

#### 通过bin/protoc自动生成rpc代码，已经写到rpcproto/compile.sh脚本。
cd rpcproto
bash -x ./compile.sh
+ ~/jungleCode/bin/protoc --go_out=plugins=grpc:. addrbookstore.proto

# ls -trl
total 24
-rw-r--r--  1 song  staff  9319 Apr  7 10:50 addrbookstore.pb.go
-rw-r--r--  1 song  staff   568 Apr  7 10:50 addrbookstore.proto

```

编译时，则`--go_out=.`变成了`--go_out=plugins=grpc:.`。



注意几个工具：

- github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-osx-x86_64.zip
  - 官网的工具：protoc用于生成protobuf的代码。
- go get -v github.com/golang/protobuf/protoc-gen-go
  - golang版本的pb插件工具，proto原型，2千多star
  - 生成的是`jungleCode/bin/protoc-gen-go`工具
  - protoc生成的xxx.pb.go代码是会引用该库的proto包。


- go get -v github.com/gogo/protobuf/protoc-gen-gogo
  - 是对golang/protobuf的分支，1千多star
  - 生成的是`jungleCode/bin/protoc-gen-gogo`工具



#### 7.2 grpc server and client

```shell
### protoBuf path

mkdir rpcserver
vim rpcserver/main.go

mkdir rpcclient
vim rpcclient/main.go

### run

```

server

```shell
# go run rpcserver/main.go
2018/04/07 18:36:18 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" > ], [phones:[number:"13812345678" ]]
2018/04/07 18:36:26 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" > ], [phones:[number:"13812345678" ]]
2018/04/07 18:36:28 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" > ], [phones:[number:"13812345678" ]]
^Csignal: interrupt
```

client

```shell
# go run rpcclient/main.go
2018/04/07 18:36:18 1

# go run rpcclient/main.go
2018/04/07 18:36:26 2

# go run rpcclient/main.go
2018/04/07 18:36:28 3
```

grpc基于http2，可以在nginx中使用。也是微服务通信的重要协议。

使用微服务，需要有较好的基础设施：

* 部署，监控，日志收集等
* 服务发现
* 服务通信，是grpc和protobuf解决的问题。
  * grpc.io
  * pb文件就是协议与接口文档，而且跨语言。
  * 调用rpc时，不需要关注底层的实现。只需要关注对端的协议和端口即可。
  * 支持服务发现。

#### 7.3 use gogo/protobuf

尝试使用github.com/gogo/protobuf库。如前面说明，这是github.com/golang/protobuf的分支。protoc默认使用的是github.com/golang/protobuf。

先尝试修改pb.go文件。

```shell
### 复制文件
mkdir rpcproto2 rpcserver2 rpcclient2
cp rpcproto/* rpcproto2/
cp rpcserver/main.go rpcserver2/
cp rpcclient/main.go rpcclient2/

# rpcproto2下的包，还是保持包名称是rpcproto
md5 rpcproto*/addrbookstore.proto
MD5 (rpcproto/addrbookstore.proto) = c13c5c0290b08eab1272329f96f2a9d8
MD5 (rpcproto2/addrbookstore.proto) = c13c5c0290b08eab1272329f96f2a9d8

cd rpcproto2
bash compile.sh
cd -
md5 rpcproto*/addrbookstore.pb.go
MD5 (rpcproto/addrbookstore.pb.go) = d4f1375e201742eaa5019be1d7bfb4d3
MD5 (rpcproto2/addrbookstore.pb.go) = d4f1375e201742eaa5019be1d7bfb4d3

## 手动修改pb.go文件
vim rpcproto2/addrbookstore.pb.go

diff rpcproto*/addrbookstore.pb.go
< import proto "github.com/golang/protobuf/proto"
---
> import proto "github.com/gogo/protobuf/proto"
24c24,25
< 	context "golang.org/x/net/context"
---
> 	"context"
>
37c38
< const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package
---
> // const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package
```

修改程序

```shell

# 修改server
vim rpcserver2/main.go
diff rpcserver*/main.go
3a4
> 	"context"
8c9
< 	"golang.org/x/net/context"
---
> 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto2"
11,12d11
<
< 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto"

# 修改client
vim  rpcclient2/main.go
diff rpcclient*/main.go
3a4
> 	"context"
6,8c7
< 	"golang.org/x/net/context"
<
< 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto"
---
> 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto2"
```

运行

```shell
# go run rpcserver2/main.go
2018/04/07 19:12:01 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" type:MOBILE > ], [phones:[number:"13812345678" type:MOBILE ]]
2018/04/07 19:12:03 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" type:MOBILE > ], [phones:[number:"13812345678" type:MOBILE ]]

# go run rpcclient2/main.go
2018/04/07 19:12:01 1

# go run rpcclient2/main.go
2018/04/07 19:12:03 2
```

效果和7.2中的结果一样。



### 7.4 proto-gen-gofast

准备环境

```shell
cp -r rpcserver2 rpcserver3
cp -r rpcclient2 rpcclient3

mkdir rpcproto3
cp -r rpcproto2/addrbookstore.proto rpcproto3

### 参考：https://github.com/gogo/protobuf
go get -v github.com/gogo/protobuf/proto
go get -v github.com/gogo/protobuf/protoc-gen-gofast

## gen
cd rpcproto3
time ~/jungleCode/bin/protoc  --gofast_out=plugins=grpc:. addrbookstore.proto
~/jungleCode/bin/protoc --gofast_out=plugins=grpc:. addrbookstore.proto  0.03s user 0.02s system 52% cpu 0.104 total

ls -lh
total 64
-rw-r--r--  1 song  staff    27K Apr  7 19:31 addrbookstore.pb.go
-rw-r--r--  1 song  staff   568B Apr  7 19:23 addrbookstore.proto

cd -
```

修改程序

```shell
vim rpcclient3/main.go
vim rpcserver3/main.go

diff rpcclient[23]/main.go
7c7
< 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto2"
---
> 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto3"

diff rpcserver[23]/main.go
9c9
< 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto2"
---
> 	"github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/rpcproto3"
```

运行

```shell
### server
go run rpcserver3/main.go
2018/04/07 19:35:46 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" > ], [phones:[number:"13812345678" ]]
2018/04/07 19:35:48 add call:[person:id:1 name:"jungle85" email:"jungle85@github.com" phones:<number:"13812345678" > ], [phones:[number:"13812345678" ]]

### client
go run rpcclient3/main.go
2018/04/07 19:35:46 1

go run rpcclient3/main.go
2018/04/07 19:35:48 2
```

需要注意的是，使用fast生成的xx.pb.go代码更长

```shell
ls -l rpcproto*/addrbookstore.pb.go
-rw-r--r--  1 song  staff   9319 Apr  7 17:38 rpcproto/addrbookstore.pb.go
-rw-r--r--  1 song  staff   9296 Apr  7 19:13 rpcproto2/addrbookstore.pb.go
-rw-r--r--  1 song  staff  27483 Apr  7 19:31 rpcproto3/addrbookstore.pb.go

```



### 8 cobra

cli工具

#### 8.1 cobra pkg

```shell
# go get -v github.com/spf13/cobra
github.com/spf13/pflag
github.com/spf13/cobra

# go get -v github.com/spf13/cobra/cobra
github.com/hashicorp/hcl/hcl/strconv
github.com/mitchellh/go-homedir
github.com/fsnotify/fsnotify
github.com/hashicorp/hcl/hcl/token
github.com/hashicorp/hcl/hcl/ast
github.com/hashicorp/hcl/hcl/scanner
github.com/hashicorp/hcl/json/token
github.com/hashicorp/hcl/json/scanner
github.com/mitchellh/mapstructure
github.com/hashicorp/hcl/hcl/parser
github.com/magiconair/properties
github.com/pelletier/go-toml
github.com/hashicorp/hcl/json/parser
github.com/hashicorp/hcl
github.com/spf13/afero/mem
github.com/spf13/afero
github.com/spf13/cast
github.com/spf13/jwalterweatherman
gopkg.in/yaml.v2
github.com/spf13/viper
github.com/spf13/cobra/cobra/cmd
github.com/spf13/cobra/cobra
```

cobra作者是spf13，写了些好用的工具，后加入google golang团队。

* [golang/go](https://github.com/golang/go) 40k star
* [gohugoio/hugo](https://github.com/gohugoio/hugo) 25k star。framework for building websites. 静态网页
* [spf13-vim](https://github.com/spf13/spf13-vim) 12k star。The ultimate vim distribution
* [cobra](https://github.com/spf13/cobra) 7k star。powerful modern CLI interfaces similar to git & go tools.
* [viper](https://github.com/spf13/viper) 5k star。golang 配置管理工具，支持json, toml, yaml等

cobra引用了一些github.com/mitchellh的包。这也是个牛人。这个人在一个高产的组织：

* https://github.com/hashicorp 

上面的cobra就引用了一些hashicorp的包。他们的明显项目：

* [vagrant](https://github.com/hashicorp/vagrant) 16k star的虚拟开发环境工具。
* [consul](https://github.com/hashicorp/consul) 



#### 8.2 生成命令框架

实现一个cli工具，用于前面的grpc的client端。

```shell
pwd
# 51reboot/day15/protoBuf/
mkdir cli
cd cli
## 下载好前面提交的工具：
# github.com/spf13/cobra
# github.com/spf13/cobra/cobra
```

生成命令工具

```shell
# ls ### cli 空目录

# cobra init
Your Cobra application is ready at
/Users/song/jungleCode/src/github.com/jungle85gopy/learn-go/51reboot/day15/protoBuf/cli.

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.

# ls
LICENSE cmd     main.go
# ls cmd
root.go

# go install
# cli
# cli help
# cli --help
#### 可以执行如上命令观察。


cobra add add
add created at cmd/add.go

# cobra add query
query created at cmd/query.go

# cobra add dump
dump created at cmd/dump.go

# ls cmd
add.go   dump.go  query.go root.go

# go install
# cli help
Usage:
  cli [command]

Available Commands:
  add         A brief description of your command
  dump        A brief description of your command
  help        Help about any command
  query       A brief description of your command

# cli help add
Usage:
  cli add [flags]

Flags:
  -h, --help   help for add

Global Flags:
      --config string   config file (default is $HOME/.cli.yaml)
```



#### 8.3 完善grpc功能

client cli

```shell
pwd
### cli

cd cmd
cp  ../../../protoBuf/rpcclient3/main.go grpc.go

vim grpc.go
# 将包名由main改成cmd
# 将main()改成newClient()函数。根据addr地址返回一个rpcproto的Client连接
# 	通过grpc.Dial获取连接conn
#	用电话本的Client封装conn，并返回
#	其余对struct的处理，转移到add.go

vim add.go
# 修改Run: func(cmd *cobra.Command, args []string)函数。
# 接收参数，并封装成一个人的个人信息，提交到grpc

```

启动服务

```shell
##### server
# cd ../../
# go run rpcserver3/main.go
2018/04/07 22:30:49 add call:[person:id:1 name:"jungle" email:"jungle@163.com" phones:<number:"13822221111" > ], [phones:[number:"13822221111" ]]
2018/04/07 22:30:52 add call:[person:id:1 name:"jungle" email:"jungle@163.com" phones:<number:"13822221112" > ], [phones:[number:"13822221112" ]]
2018/04/07 22:30:54 add call:[person:id:1 name:"jungle" email:"jungle@163.com" phones:<number:"13822221113" > ], [phones:[number:"13822221113" ]]


##### client at day15/protoBuf/cli
# pwd

# go build
# ./cli add 1 jungle jungle@163.com 13822221111
add called, args: [1 jungle jungle@163.com 13822221111]
add ok, id: 1

# ./cli add 1 jungle jungle@163.com 13822221112
add called, args: [1 jungle jungle@163.com 13822221112]
add ok, id: 2

# ./cli add 1 jungle jungle@163.com 13822221113
add called, args: [1 jungle jungle@163.com 13822221113]
add ok, id: 3
```

