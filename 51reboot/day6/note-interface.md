

## 接口

为什么需要接口。规一化一类东西。比如usb接口。不管你是电脑还是手机，或者是电视、音响。符合某个确定的规范。不管来自什么厂家。大家都可以使用该规格的接口互联。



### 1 文件相关的接口。

go定义的一些接口：

io相关


```Go
package io

type Writer interface {
    Write(p []byte) (n int, err error)
}
// 继续使用type来做定义，这次不再是struct，而是interface
```

文件相关的接口


```go
package os

func (f *File) Write(b []byte) (n int, err error)
// os.*File 类型的方法Write即实现了io.Writer接口。
```

格式化写入接口




```go
package fmt

func Fprintf(w io.Writer, format string, a ...interface{})
(n int, err error)
// 将fmt的内容写入到io.Writer这个接口的实例w中。
// 这里w是一个文件接口。
```


实现了接口里所有的方法，即实现了该接口。



### 2 实例

* interface-test-1.go 实现距离接口，包括平面上点的距离、折线的长度。
* ByteCounter-struct-2.go Write接口实现基于结构体的字节计数器。
* LineCounter-3.go 新建类型通过Write接口实现的行计数器。
* ByteCounter-interface-4.go 新建类型通过Write接口实现字节计数
* bytes-buffer-5.go 通过bytes.Buffer做为输入来统计行和字节。
* multiWriter/main.go 通过MultiWriter接口实现统计行数和统计字节数。
* tee/main.go 通过teeReader接口实现统计行数和统计字节数。



#### 模拟http monitor统计

```shell
# run
go run http-monitor/main.go &

# access 
curl localhost:9090
hello Golang

curl localhost:9090/sdf
hello Golang

curl localhost:9090/monitor
counter:21
curl localhost:9090/monitor
counter:25
```



#### errors

实现了Error()方法，即实现了错误处理接口。

```go
var e error
e = errors.New("an error")
fmt.Println(e.Error())
fmt.Println(e)
// an error
// an error

e = fmt.Errorf("err from fmt.Errorf")
fmt.Println(e)
// err from fmt.Errorf
```





    ls -l 
    -rw-r--r--  1 song  staff  1114 Jul  8 14:31 interface.go
    -rw-r--r--  1 song  staff     6 Jul  8 14:31 a.txt
    -rw-r--r--  1 song  staff   443 Jul  8 14:55 ByteCounter-struct.go
    -rw-r--r--  1 song  staff   574 Jul  8 15:21 LineCounter.go
    -rw-r--r--  1 song  staff   454 Jul  8 15:22 ByteCounter-interface.go
    drwxr-xr-x  3 song  staff   102 Jul  8 15:28 multi
    drwxr-xr-x  3 song  staff   102 Jul  8 15:29 tee
    drwxr-xr-x  3 song  staff   102 Jul  8 15:54 http-monitor
    -rw-r--r--  1 song  staff   283 Jul  8 16:06 errors.go



#### 接口断言

用于确定接口的类型。

```go
f := w.(*os.File)
c := w.(*bytes.Buffer)
// 断言w的类型，括号后是希望断言的类型
```



多个接口断言，可以用switch

```go
  var x interface{}
  swich x.(type) {
      case nil: 
      case int:
      case bool:
      case string:
      default:
  }
```



#### 空接口

任意类型都实现了空接口。因此空接口可以接受任何类型。

blankInferface.go 给出了使用示例。将int, string, Point都赋值给空接口。

```go
type Blank interface {}
// 任何类型都实现了该接口。

// 一些使用空接口的例子
Marshal(v interface{}) ([]byte, error)

Unmarshal(data []byte, v interface{}) error

fmt.Println(a ...interface{}) (n int, err error)
```





