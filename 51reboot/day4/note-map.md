
### map

- hash方式的
- 无序的
- O(1)的访问时间



```go
ages1 := make(map[string]int)
// map的定义方式

// 赋值方式有2种，定义时同时赋值，使用{}给出；另一种是先定义，再单独按项赋值。

// 对map的取值。通常是如下方式。根据ok的值来判定结果。
v, ok := ages1["c"]
// 如果map中有c，则将值赋给v，同时ok为true
// 如果map中无c，则value的值为对应类型的空值，ok为false
```

对map的遍历，直接使用for … range。但有2种方式。

```go
// key, value
for name, age := range ages {
    fmt.Println("name", name, "age", age)
}
// range key
for name := range ages {
    fmt.Println(name)
}
```

Code: map/main.go



删除操作：

```go
// 删除map中key所对应的value.
delete(dict, 'key')
```

#### 词频统计

通过map来保存词频。

如上述所说的判断一个key是否在map中，如果不在，则返回的值为0。因此下面的2段代码效果是相同的。这是默认值带来的简洁。

```go
// style 1
counter[word]++

// style 2
if _, ok := counter[word]; ok {
    counter[word]++
} else {
    counter[word] = 1
}
```

效果：

```shell
cat a.txt
   an     apple a qq a golang

   a python

go run counter/main.go a.txt | sort -k1
a	3
an	1
apple	1
golang	1
python	1
qq	1
```

code: counter/main.go



#### set

code:mapSet/main.go



