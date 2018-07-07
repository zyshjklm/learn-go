

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



main6.go 获取字段名

```shell
# go run main6.go
main.Http
name:{Name:name PkgPath:main Type:string Tag: Offset:0 Index:[2 0] Anonymous:false}
type:string

password:{Name:password PkgPath:main Type:string Tag: Offset:16 Index:[1] Anonymous:false}
type:string

```

获取字段的2种方式：

* t.FieldByName() 基于字段名，包括匿名嵌套的字段名。
* t.FieldByIndex() 基于字段索引。 

注意运行结果中，name的索引Index值，与password值中的Index的差别。



main7.go 获取结构体的方法

```shell
# go run main7.go
*main.Http
{Name:GetAgent PkgPath: Type:func(*main.Http) string Func:<func(*main.Http) string Value> Index:0}
{Name:GetHost PkgPath: Type:func(*main.Http) string Func:<func(*main.Http) string Value> Index:1}
{Name:GetName PkgPath: Type:func(*main.Http) string Func:<func(*main.Http) string Value> Index:2}
{Name:GetPass PkgPath: Type:func(*main.Http) string Func:<func(*main.Http) string Value> Index:3}
```

注意点：

* 需要使用指针变量的反射来操作
* 方法的数量使用t.NumMethod()函数。
* t.Method(i) 返回的是第i个方法结构体。其主要字段：
  * Name 方法名
  * PkgPath 包路径，即包是在那里定义的
  * Func 即函数的声明方式，包括参数及返回类型
  * Index 索引值



main8.go 结构体方法的详细比较之一：

* Http及Data结构体都各定义2个方法，且都使用指针接受者。
* 反射时使用指针，但在获取方法时，分别对比指针及其Elem。

```shell
# go run main8.go
ptr: *main.Http
{GetAgent  func(*main.Http) string <func(*main.Http) string Value> 0}
{GetHost  func(*main.Http) string <func(*main.Http) string Value> 1}
{GetName  func(*main.Http) string <func(*main.Http) string Value> 2}
{GetPass  func(*main.Http) string <func(*main.Http) string Value> 3}
elem: main.Http
```

从上可见，通过指针反射，可能获取到所有使用指针接受者的方法，使用Elem()变量却不能获取指针接受者方法。



main9.go 主要是与main8.go形成对比。

改变了2个方法的接受者类型。

```shell
# go run main9.go
ptr: *main.Http
{GetAgent  func(*main.Http) string <func(*main.Http) string Value> 0}
{GetHost  func(*main.Http) string <func(*main.Http) string Value> 1}
{GetName  func(*main.Http) string <func(*main.Http) string Value> 2}
{GetPass  func(*main.Http) string <func(*main.Http) string Value> 3}
elem: main.Http
{GetAgent  func(main.Http) string <func(main.Http) string Value> 0}
{GetPass  func(main.Http) string <func(main.Http) string Value> 1}
```

当GetAgent及GetPass使用变量为接受者时：

* 指针反射可以获取到所有类型的方法，包括指针接受者及变量接受者。
* 变量反射只能获取到使用变量为接受者的方法。

反射是建立在数据类型的基础之上的。需要是静态类型语言。



### 结构体tag

结构体的tag存在坑。需要特别注意。

```shell
# cd ../
# mkdir reflecttag
# cd reflecttag

# go run mainTag1.go
{"DiskIOPS":"100"}
d:{DiskIOPS:100}
s:{DiskIOPS:}

# go run mainTag2.go
{"DiskIOPS":"100"}
d:{DiskIOPS:100}
s:{DiskIOPS:100}

# diff mainTag1.go mainTag2.go
9c9
< 	DiskIOPS string `json:"disk-IOPS"`
---
> 	DiskIOPS string `json:"diskiops"`
28d27
<
```

这两个程序唯一的差别就是tag的名称有点差异。mainTag2.go中，tag与字段DiskIOPS只有大小写的差别。因此可以反序列化。其间的逻辑关系：

* 对d进行序列化，而D没有定义json序列化的tag，因此使用了字段名DiskIOPS。
* 故序列化的结果，jd使用DiskIOPS作为Key.
* 将jd反序列化到s时，S设置了json的tag为diskiops，这与jd中的DiskIOPS只有大小写差异
* 因此可以反序列化，且使用S的DiskIOPS为字段名。