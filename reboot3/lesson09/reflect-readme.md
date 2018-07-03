

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



