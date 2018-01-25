
# lesson3 



## 1 整数

字节数：

type/type.go

```shell
int 	 8
int32 	 4
int64 	 8
uint 	 8
uint32 	 4
uint64 	 8

int8 	 1
uint8 	 1
```

这是64位系统的结果。这里用到了reflect和unsafe包。

```go
fmt.Println(reflect.TypeOf(x1), "\t", unsafe.Sizeof(x1))
```

int后面的数字，代表的是**byte数**。

一定要注意变量的字节情况，特别是写网络程序，收发包在双方的系统差异大时，不注意可能出现问题。溢出，字节序等。

变量的范围

uint32的范围

```python
>>> 2 ** 32 - 1
4294967295
>>> pow(2, 32) -1
4294967295
# 0 - 4294967295

# int32
>>> pow(2, 31) -1
2147483647
>>> pow(2, 31)
2147483648

# (-2^16)-(2^16-1)
# - 2147483648  -  2147483647
```

int8和uint8都使用1个字节。


## string

str/string.go

输出字符使用"%c"
字符串不可修改，通过[]byte的转换来完成修改。

xxd main.go

### 浮点数

float/float.go

```go
+
-
*
/
%	//取模

if (a > b && b > 3) ||
```

## 常量 

const/const.go

const

iota 自动加1.

## 强类型

类型转换需要显式进行。

字符串跟数字需要借助函数进行转换。strconv

```python
'0x' % (1024129)
'fa081'
# int向uint8转换。
```



使用strconv.FormatInt生成随机串。

代码：strong/main.go



## if 

```go

var x bool
if x {
	fmt.Println("true")
} else {	
	fmt.Println("false")
}
```

if 没有小括号；只能是bool表达式，数字。字符串之类的不行。

if/main.go

