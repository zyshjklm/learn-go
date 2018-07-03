

## 反射reflect

### 基本用法

main1.go 通过反射获取类型：

```shell
# mkdir reflect
# vim main1.go
# go run main1.go
X32 int32
```

如下：

- Name() 是自定义的类型
- Kind() 是底层的类型。



main2.go 获取变量及其指针的相关类型

```shell
# go run main2.go
var:
	Name: X32, Kind: int32
pointer:
	Kind: ptr, Elem: main.X32
```

与前面的变量相比，指针对应的类型分别是：

* Kind() 是类型，此处是`ptr`值。
* Elem() 指向具体的元素，是对应变量的Name()



main3.go 获取变量的值

```shell
# go run main3.go
type: main.X32, value:100
```

通过ValueOf获取原变量的反射值v：

* v 是原变量的值
* v.Type()是原变量的类型，对应反射类型的Kind()



### 结构体的反射

main4.go 先定义了一个嵌套的结构体。再通过反射来获取其类型。并分别操作结构体变量，结构体指针的反射。

```shell
# go run main4.go
*main.Http
main.Http
main.Http
```

其中`reflect.Ptr`是个变量，表示的是反射指针。比较时需要使用Kind()。



main5.go 通过反射遍历结构体的字段Field。包括遍历嵌套字段。

```shell
# go run main5.go
*main.Http
type:string,name:host,index:[0],offset:0
type:string,name:agent,index:[1],offset:16
type:main.data,name:data,index:[2],offset:32
-- type:string,name:name,index:[0],offset:0
-- type:string,name:password,index:[1],offset:16
```

最后2行即是嵌套的data字段。

主要知识点：

* t.NumField() 是字段的数量
* t.Field(i) 是获取第i个字段的描述结构体。包括Type, Name, Index, Offset等
* 描述结构体的`f.Anonymous`代表是否匿名字段
* 获取嵌套字段使用上述的`Type`变量。

