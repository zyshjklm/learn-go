



## json encode



refer: https://www.jianshu.com/p/f3c2105bd06b



### base1marshal.go struct字段可见性

结构体中password是不可见的。因此编码为时json没有该字段。



### base2complex.go 复合结构

主要是切片与map。map的key必须是string，值为同一类型的数据。



### base3embed.go 嵌套结构体

将前2种方式嵌套使用。因为map的value需要是同一类型，而通过嵌套结构体，由结构体来实现不同类型的key和value。



### base4diffArray.go 类型不统一数组

基于base2complex.go，增加一个类型不统一的数组字段。需要通过空接口`[]interface{}`来定义值的类型。



### base5mapInterface.go 异构map值

使用空接口`[]interface{}`来定义map值的类型。

空接口定义的零值是 nil。编码成json是null。本示例中相应的结果：

```json
"Extra":null,
"Level":{"server":90,"tool":null,"web":"Good"}
```



### tag1Base.go tag基本用法

golang的struct需要字段首字母大写才能导出，而json的世界更盛行小写字母的方式。解决这对矛盾的办法就是使用tag来重命名结构字段的输出形式。



### tag2Ignore.go 忽略可见字段

通过`-`占位来忽略可见字段。比如Password字段是对外可见了，但又不希望编码成json时输出该字段。

```go
type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Money    float64 `json:"money"`
}
// 忽略Password字段
```



### tag3Omit.go omitempty

通过设置omitempty，当字段的值为零值时，则不输出该字段。

```shell
go run tag3Omit.go
{"email":"rsj217@gmail.com","money":100.5}
// password 字段的值为空串

{"email":"rsj217@gmail.com","password":"123456","money":100.5}
// password有值
```



### tag4String.go 转换类型

将结构体中的数据转换成json的字符串。

```go
// struct
Money    float64 `json:"money,string"`

// 编码成json时，变成了 money: "100.5"
```

反之，通过这样的定义，也可以将json中的数字串，转换成结构体中的数字。

代码实现参考：tag5Unmarshal.go