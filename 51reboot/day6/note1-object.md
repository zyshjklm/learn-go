## 对象

问题：函数回顾。

* 学生系统，如果增加一个班级的概念，应该如何处理？
* 两点之间的距离，函数式的使用结构体表示点。再用一个函数来处理2个结构体。



面向对象：声明方法。

鸭子模型。

type 定义新的类型。

我们不能直接对对系统的类型直接声明方法。如对string类型增加或者改变方法。但我们可以通过 type来对string定义新的类型，如 type MyString string。再对MyString声明新的方法。

匿名类型

```go
var anonimus struct {
  Id: int
  Name: string
}
anonimus 
```

注意下面的区别

```go
var  p Path	// variable 变量
type q Path	// type 类型
```



学习两种不同的实现方法：

* 基于结构体的函数
* 基于类型的方法。



#### 函数到方法

* method-point-dist.go。object/method1-dist-point.go。基于Point结构体
  * 实现其计算距离的Distance(p, q Point)函数。
  * 实现计算距离的方法p.Distance(q Point)
* method-pointSlice-dist.go。object/method2-dist-slice.go。基于[]Point
  * 实现Point的Distance方法
  * 实现[]Point的距离函数Distance(  path []Point) (length float64)
  * 实现[]Point的距离函数Distance1(path []Point) (length float64)
  * 这里没有[]Point的距离方法。
* method-path-dist.go。object/method3-dist-path.go。基于Path类型，代替[]Point。
  * 实现Point的Distance方法
  * 定义Path类型： type Path []Point
  * 实现Path的距离方法 (path Path) Distance() float64 
* method-pointer.go 基于Point说明指针及方法的定义与用法
  * 定义时使用指针。(p *Point) ScaleBy(factor float64)
  * 使用时不显示使用指针：
    * p3 := Point{1, 2}
    * p3.ScaleBy(4)



注意如下几种方式的区别

```go
func (p Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func ScaleBy(p Point, factor float64) {
	p.X *= factor
	p.Y *= factor
}

func ScaleBy((p *Point, factor float64) {
	p.X *= factor
	p.Y *= factor
}
```

(p Point)是值拷贝，不会影响p的内容。

(p *Point)是指针。能影响p的内容。

更详细的内容参考：http://studygolang.com/articles/4059



### 可见性

* 通过首字母大小来控制可见性
* 可见性是package级别的
  * 在package内部，可以访问所有变量和方法。
  * 在package外部，只能访问大写开始的变量和方法。

