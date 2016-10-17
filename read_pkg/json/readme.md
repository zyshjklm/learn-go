
## json package

read : https://golang.org/pkg/encoding/json/

### 1 introduction of json

see: https://blog.golang.org/json-and-go

#### 1.1 encode and decode

* encoding struct to json
* decoding json to struct
* decoding []byte to struct

#### 1.2 generic json and decoding

the default concrete type of Go:

* bool 
* float64
* string
* nil

the json package use map[string] interface{} and []interface{}
values to store arbitrary JSON object and arrays.

for arbitrary decoding. json supports reference types.
* pointer
* slice
* map



#### 1.3 streaming 

* Decoder, Encoder for readming and writing streams of JSON data.
* NewDecoder, NewEncoder wraps the io.Reader, io.Writer interface.

due to the ubiquity of Readers and Writers, the New* func can be used in a board range of scenarios, such as reading and writing to HTTP connections, WebSockets, files.



example for Newdecoder and NewEncoder. 用于提取想要的字段。

```shell
./json_stream >out
# input here:

{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}
{"Name":"Jungle"}
{}
{"Age":6}
abd
2016/10/17 12:18:46 invalid character 'a' looking for beginning of value
```

当输入abd并回车后，程序退出。

对就的。out的值为：

```json
{"Name":"Wednesday"}
{"Name":"Jungle"}
{}
{}
```

### 2 pkg/encoding/json

read : https://golang.org/pkg/encoding/json/

```go
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error

// usage:
var out bytes.Buffer
json.Indent(&out, b, "=", "\t")
out.WriteTo(os.Stdout)
```

Indent appends to dst an indented form of the JSON-encoded src. Each element in a JSON object or array begins on a new, indented line beginning with prefix followed by one or more copies of indent according to the indentation nesting.

使用时需要使用bytes包中的bytes.Buffer类型。然后使用其WriteTo方法输出。



