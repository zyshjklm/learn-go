
注意变量声明

```go
var students = make(map[string]Student)
// 适合所有地方使用，特别是全局变量

students := make(map[string]Student)
// 仅适合在函数里面。不能用于全局变量。

```





### leetcode



https://leetcode-cn.com/



https://github.com/golang/groupcache/blob/master/lru/lru.go



有个工作场景求指点思路。mongo中存了一些文档。但会不断增加新的字段，如果用golang处理这些mongo文档，需要事先定义好对应的strcut，每增加一个字段就要改struct比较麻烦，有没有办法，根据读到的mongo文档字段，再动态的构造一个结构体。或者用一个文件预先存储好mongo中字段的名字与类型，根据这个文件构造出struct。

https://blog.csdn.net/pkueecser/article/details/50412731


