## lector slice

```go
/*
	slice 并不是数组，是引用，而是指向底层的数组，
  	slice做为变长数组的替代方案。
    可以直接创建，或者从数组生成。
    make([T], len, cap), len(), cap()

	通过cap可以达到变长数组的效果。对数组增加元素时，如果还有容量，则不影响数组长度
	通过起过容量，则新分配一个原数组两倍长的数组。

	append操作：
	可以在slice尾部追加元素
    也可以将一个slice追加到另一个slice尾部
    如果追加后的最终长度未起过 追加到的slice的容量，则返回原始slice
    否则，将重新分配数组，并拷贝原始数据

*/

```