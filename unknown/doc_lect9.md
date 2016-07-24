## lect9 function

go函数**不支持**：

* 嵌套
* 重载
* 默认参数

支持如下特性：

* 无需声明原型
* 不定长参数 
* 多返回值
* 命名返回值参数
* 匿名函数
* 闭包



### closure

##### a closure is a [record](https://en.wikipedia.org/wiki/Record_(computer_science)) storing a function together with an environment:

##### a mapping associating each [free variable](https://en.wikipedia.org/wiki/Free_variable) of the function (variables that are used locally, but defined in an enclosing scope) with the [value](https://en.wikipedia.org/wiki/Value_(computer_science)) or [reference](https://en.wikipedia.org/wiki/Reference_(computer_science)) to which the name was bound when the closure was created.



```go

func closureOut(x int) func(int) int {
	return func(y int) int {
        return x + y
	}
}
```



### defer

defer的执行方式类似其它语言的析构函数。在函数执行结束时，按调用顺序的相反顺序依次执行。

即使函数发生严重错误时也会执行。

支持匿名函数的调用

学用于资源清理，文件关闭，解锁，记录时间等

与匿名函数配合，可以return之后修改函数计算结果。

