### 格式化

fmt 

```shell
gofmt -w hello.go
```

#### 要求

* 编译通过
* gofmt统一风格
* 代码不提交二进制文件



Note: python的小整数都是一个内存引用地址。(进行了优化）

```python
python
Python 2.7.13 (default, Dec 18 2016, 07:03:39)
[GCC 4.2.1 Compatible Apple LLVM 8.0.0 (clang-800.0.42.1)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> a=1
>>> b=1
>>> id(a)
140685357395576
>>> id(b)
140685357395576

```



### 自举

使用语言自己来写自己。

* c开始是用的其它语言，之后再用自己写自己。
* go从1.5开始就自举。



GOROOT 二进制

GOPATH 



### myEchoArgs1

```shell
ls myEchoArgs1
main.go sep.go

go run myEchoArgs1/main.go
# command-line-arguments
myEchoArgs1/main.go:15:8: undefined: sep
myEchoArgs1/main.go:17:3: undefined: sep

go run myEchoArgs1/main.go myEchoArgs1/sep.go true x yy zzz golang
[/var/folders/4_/xxx/T/go-build593391458/command-line-arguments/_obj/exe/main true x yy zzz golang]
true
arg line:true x yy zzz golang

## 说明：main.go sep.go都声明为package main。故在run时需要全部指定

cd myEchoArgs1
go build
./myEchoArgs1 true x yy zzz golang
[./myEchoArgs1 true x yy zzz golang]
true
arg line:true x yy zzz golang
```



### myEchoArgs2

```shell

cd myEchoArgs2
go build 
# 生成 myEchoArgs2
go build main.go
# 生成 main

./main -h
Usage of ./main:
  -n	end with
	 for string
  -s string
    	separator (default " ")

# 注意参数-n的差别
./main x yy zzz golang; echo ":over"
x yy zzz golang:over

./main -n x yy zzz golang; echo ":over"
x yy zzz golang
:over

/main -s "---" x yy zzz golang; echo ":over"
x---yy---zzz---golang:over

./main -n -s "---" x yy zzz golang; echo ":over"
x---yy---zzz---golang
:over
```



#### flag

```go
flag.Bool("n", false, "end with new line")
flag.String("s", " ", "separator")
// 通过命令行携带参数。但这种方式设置的参数，不能通过-h的方式查看。

flag.StringVar(&myType, "t", "test", "test, prod, preview")
// 可以通过 -h查看，包括前面的String函数设置的flag
```



### godoc

#### godoc.org

用于在线查看文档。

* godoc.org
* godoc.org/fmt
* godoc.org/github.com/go-redis/redis
* godoc.org/github.com/icexin/golib



#### 本地doc



```shell
godoc -http ":6060"

http://localhost:6060/pkg/os/
```



go的packge本质上讲，是文件的路径。因为package的文件名的上级路径，如github.com/go-redis/redis。



### go常用命令

```shell
go install golang.org/x/tools/cmd/godoc
```



* run 针对一个文件，而不是package
* build 针对package
* install 针对package



gopath的src目录下。

其全路径为src下的路径，import path。




### godoc



### 变量：

零值初始化。初始化为零值，是为了安全。

短变量。通常用于局部变量。

```go
i := 0	// int
s := 'hello'	// string
i, j := 0, 1 	// bat init
```



未使用的变量，引入而未使用的包，会导致出错。以此来禁止滥用包。

一方面是package有依赖的检查，增加编译时间；另一方面有二进制包的大小增加。



## 指针

* *T 类型为T的指针，` *int`, ` *float32`,
* &t 取变量t的地址，可用于赋值给指针变量
* *p 即取指针变量所指向的内容。



*int 是指针，&x 是指针

指针的值不能直接进行修改？

缓冲区溢出。C可以指向任何一块内存。



### 变量的生命周期

垃圾回收

栈与堆



生命周期

for i := 1; i < len(os.Args); i++ {

}

栈：用于临时变量；函数开始时申请，函数结束时释放。

堆：用于动态管理。程序申请时开辟，程序结束时释放。



原本是数据结构的概念；栈是先进后出；堆则先进先出。

```go
func localFunc() {
  var local = "local"
  fmt.Println(local)
}
// local 是在栈上分配的。

func main() {
  p := localFunc()
  fmt.Println(*p)
}

// 逃逸分析。
// 因为要返回local变量的地址，因此原本在栈上分配的变量，转移到在堆上分析了
func localFunc() *string {
  var local = "local"
  fmt.Println(local)
  return &local
}

func main() {
  p := localFunc()
  fmt.Println(*p)
}
```

上面这种方式，对C语言来说，则会导致数据内容不确定。因为局部变量的地址空间已经释放了。

### 多变量赋值

变量类型显示声明，现在版本的php, python等动态语言，还在加入类型。因为项目规模很大时，动态语言的无类型，可能导致相关的困扰。另外还也内存空间的使用量及性能问题。

```go

x, y = y, x

func fib(n int) int {
  for i := 0; i < n; i++ {
    //
  }
}
```



## 作用域



爬虫

广度优先与深度优先。

广度优先是一个同心圆向外扩。



## 作业

函数里面使用一个变量时，需要进行查询。语句内部，函数内部，全局变量。

