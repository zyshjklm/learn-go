## 复合数据类型

### 1 数组

```go
var q [3]int = [3]int{1,2,3}

// 只赋值了2个元素
var r [3]int = [3]int{1,2}

// 指定索引位置进行赋值
q2 := [...]int{4: 2, 10: -1}
```

code: arrary/array.go



### 2 练习md5

直接使用了crypto/md5包中的Sum()函数。

如下结果显示与操作系统的工具结果相同。

```Shell
go build md5/main.go
./main rune.go note-homework.md
c4b477caf7ac883eaa7f6b3a33878148 rune.go
9364a2e5cd78c1f8969471dc9d8c6e24 note-homework.md

md5 rune.go note-homework.md
MD5 (rune.go) = c4b477caf7ac883eaa7f6b3a33878148
MD5 (note-homework.md) = 9364a2e5cd78c1f8969471dc9d8c6e24
```

code: md5/main.go



### 3 slice

既然有了数组，为什么还要有slice。因为数组是定长的。有时我们希望能处理变长的情况，而且最好的不需要关注变长的细节。

```go
primes := [6]int{2, 3, 5, 7, 11, 13}

// 获取了primes的类型为 [6]int，是数组
fmt.Println("type of primes:", reflect.TypeOf(primes))

var s []int = primes[1:4]

// 获取了s的类型为 []int，是slice
fmt.Println("type of s:", reflect.TypeOf(s))
fmt.Println(s)

// 0xc4200720c8
fmt.Println(&s[0])

// 0xc4200720c8 . 两者的地址相同。
fmt.Println(&primes[1])

```



注意对切片的多次切片，以及通过子切片来修改原切片某元素的值。

#### 空切片 nil

空切片的len, cap都是0，其指向的内存空间为nil。其本身的值也等于nil

```go
// nil
var s3 []int
fmt.Println(s3, len(s3), cap(s3))
if s3 == nil {
    fmt.Println("nil!")
}
// [] 0 0
// nil!
```

code: slice/main.go



#### append slice

通过在slice中增加元素，来观察其len和cap。

可以看到cap是按2的指数幂增长的。

```shell
[]
len=1 cap=1 [1]
len=2 cap=2 [1 2]
len=3 cap=4 [1 2 3]
len=4 cap=4 [1 2 3 4]
len=5 cap=8 [1 2 3 4 5]
len=6 cap=8 [1 2 3 4 5 6]
len=7 cap=8 [1 2 3 4 5 6 7]
len=8 cap=8 [1 2 3 4 5 6 7 8]
len=9 cap=16 [1 2 3 4 5 6 7 8 9]
len=10 cap=16 [1 2 3 4 5 6 7 8 9 10]
len=11 cap=16 [1 2 3 4 5 6 7 8 9 10 11]
len=12 cap=16 [1 2 3 4 5 6 7 8 9 10 11 12]
len=13 cap=16 [1 2 3 4 5 6 7 8 9 10 11 12 13]
len=14 cap=16 [1 2 3 4 5 6 7 8 9 10 11 12 13 14]
len=15 cap=16 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
len=16 cap=16 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
len=17 cap=32 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]
```

Code: append/main.go

