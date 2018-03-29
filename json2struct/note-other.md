

## 其它



### 双重数组的解码

有两种办法

* 通过json-to-go解析成双重数组。如doubleSlice1.go
  * 示例：**Asks \[][]int \`json:"Asks"`**
  * 底层直接直接是int型。可以直接解码。
* 先定义一个结构体，再用该结构做为另一个结构的`[]`成员类型。如doubleSlice2.go
  * 示例：**Asks []Order\`json:"Bids"`**;  Order本身又是一个结构
  * Order需要解码函数`func (o *Order) UnmarshalJSON(data []byte) error {}`

第二种方法处理得更麻烦。

```shell
#go run doubleSlice1.go
{Asks:[[21 1] [22 1]] Bids:[[20 1] [19 1]]}

#go run doubleSlice2.go
{Asks:[{Price:20 Volume:1} {Price:19 Volume:1}] Bids:[{Price:21 Volume:1} {Price:22 Volume:1}]}

```

