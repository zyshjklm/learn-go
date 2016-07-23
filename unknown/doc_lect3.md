## 3 type and variable


```go
// 常量
const (
    PI = 3.14
    const1 = "1"
    const2 = 2
    const3 = 3
)

// 变量
var (
    name = "gopher"
    name1 = "1"
    name2 = 2
)

// 自定义类型
type (
	newType int
	type1 float32
	type2 string
)
```



### 常用类型

基本类型：

- 布尔: bool: true, false
- 整型：
  -  int/uint 根据平台确定是32，64
  - 8位：int8/uint8
  - 16位: int16/uint16
  - 字节型：byte(uint8的别名)
  - 32：int32(rune)/uint32
  - 64:	int64/uint64

- 浮点型：
  -  float32, float64

- 复数：
  -  complex64/complex128

- 足够保存指针的32/64位整型：
  - uintptr

- 其它：
  -  array, struct, string

- 引用：
  - slice, map, chan(channel)

- 接口类型：
  -  interface

- 函数类型：
  - func，函数也可以做为变量值 


### 零值 

零值：

	通常是0
	string: ""
	bool: false

### 全局变量

var 全局变量，且可以并行声明
var a, b, c, d = 1, 3, 5, 7
or 
a, b, c, d := 1, 3, 5, 7

:=是用来代替var的，表示是全局变量
但在var ()中就不能再使用:=赋值了。

类型转换：

<ValueA> [:]= <TypeOfValueA> (<ValueB>)
