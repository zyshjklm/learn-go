

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



### 不可见字段的编码

不可见字段是不可访问的，也不能直接被编码到json中。因此需要自行实现MarshalJSON，从而方便编码。

额外的，也实现了UnmarshalJSON，只不过两者在实现时有差异：

- MarshalJSON直接使用map定义一个临时的对象。
- UnmarshalJSON需要定义一个可导出的临时结构体用于内部Unmarshal

关于UnmarshalJSON时定义结构体的问题，可以通过https://github.com/json-iterator/go的私有字段来解决。

```shell
# go run other/unimported.go
2018/03/29 00:11:05 {"age":23,"id":1,"name":"jungle85"}
2018/03/29 00:11:05 &main.Student{stuID:1, stuName:"jungle85", stuAge:23}
```



### json包含转义字符

```json
{
   "id": 12345,
   "name": "Test Document",
   "payload": "{\"message\":\"hello!\"}"
}
```

payload值是字符串，而且有转义字符。需要解析出这个串中的key/value。

处理思路：

* 先定义两级结构体进行嵌套。
  * 这里只能解析到2层结构体，却得不到里层的数据。且会报错：
  * cannot unmarshal string into .. field LogEntry.payload of type main.LogPayload
* 实现LogPayload的UnmarshalJSON
  * 先从[]byte中解析成字符串，
  * 再从串中解析到payLoad结构体。这里用到一个技巧：通过type定义一个临时类型
  * 通过fakeLogPayload得到结果，再转换到LogPayload

运行结果：

```shell
#go run escapeJSON.go
main.LogEntry{ID:12345, Name:"Test Document", Payload:main.LogPayload{Message:"test"}}
main.LogPayload{Message:"test"}
```

