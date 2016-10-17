
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



