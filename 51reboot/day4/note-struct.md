### struct

结构体是一种新的数据类型。是对一类事物多方面属性的统一描述。

通过type structName struct 来定义。需要注意几种类型的字节空间占用。

```go
var s Student
s.Id = 1
s.Name = "jack where are you going to fly! "
var str = s.Name
fmt.Println(unsafe.Sizeof(s))      // 24
fmt.Println(unsafe.Sizeof(s.Id))   // 8
fmt.Println(unsafe.Sizeof(s.Name)) // 16
fmt.Println(unsafe.Sizeof(str))    // 16
fmt.Println(len(str))              // 3
// 上述的单位是字节数。uint8的长度是1字节。

// 字符串变量str，变量其实是一个指针，占用16字节。
// 而str指向的空间的长度是3字节。
```

另外，结构体本身可以通过指针来标识。当指针p指向一个结构体时，如上的s变量，则p.Name和s.Name是一个变量。

