## lect8 map



```go
/*
map以key:value的形式存储
查找速度：
	索引比map快2个量级，map比线性搜索快
使用make创建map:
*/

make(map[int]string)
// key -> int
// value -> string

// 多重嵌套时，需要从外层逐步向里层声明
m3 = make(map[int]map[int]string)
m3[1] = make(map[int]string)
m3[1][1] = "OK"
	
```

