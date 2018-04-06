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

