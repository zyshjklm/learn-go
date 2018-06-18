## go by example



### 00 hello

```shell
# go run 00-hello.go
hello golang!

# ls
00-hello.go readme.md

# go build
# ls
00-hello.go byExample   readme.md

# ./byExample
hello golang!

```



### 01 values

```SHELL
# go run 01-values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```



### 02 variables

```shell
# go run 02-variables.go
initial
1 2
true
0
short
```



### 03 constant

```shell
# go run 03-constants.go
constant
6e+11
600000000000
-0.28470407323754404
```



### 04 for

`for` is the only one loop construct with multiple styles.

```shell
# go run 04-for.go
1
2
3
7
8
9
loop
1
3
5
```



### 05 if-else

```shell
# go run 05-if-else.go
7 is odd
8 is divisible by 4
9 has 1 digit
```



### 06 switch

```shell
# go run 06-switch.go
Write 2 as two
it's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
```



### 07 array

array: [3]int. TYPE and LEN

```shell
# go run 07-arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d: [[0 1 2] [1 2 3]]
```



### 08 slice

slice are a key **data type** IN GO, giving a more powerful interface to sequences than arrsys. 

slice are typed only by the elements they contain.

`make([]string, 3)` create a slice of strings of length 3.

```shell
# go run 08-slice.go
emp [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d: [[0] [1 2] [2 3 4]]
```



#### 08-1 array and slice

refer: https://blog.golang.org/go-slices-usage-and-internals

**Arrays**

An array type definition specifies a length and an element type. [4]int represents an array of for integers. An array's size is fixed; its length is part of its type.

the zero value of an array is a ready-to-use array whose elements are themselves zeroed.

important: Go's array are values. An array variable denotes the entire array; it is not a pointer to the first array element as would be the case in C.

**Slices**

Slices are build on arrays to provide great power and convienience. Unlike an array type, a slice type has no specified length.

```go
# array
arr := [5]int

# slice
letters := []string{"a", "b", "c"}
```

slice can be created with the make func: `func make([]T, len, cap) []T`.

*   len: length of []T slice
*   cap: optional capacity. 
    *   when the cap argument is omitted, it defaults to the specified length.
*   make allocates an array and return a slice that refers to the array.

```go
s := make([]byte, 5)
len(s) == 5
cap(s) == 5
```

the zero value of a slice is `nil`. 

**len and cap**

Slicing is done by specifying a half-open range with two indices separated by a colon.

**Slice internals**

A slice is a descriptor of an array segment. It consists of:

*   a pointer to the array
*   the length of the segment
*   capacity(the maximum length of the segment)

example:

`slice1 := make([]byte, 5)`

![slice1](https://blog.golang.org/go-slices-usage-and-internals_slice-1.png)

`s = slice1[2:4]`

![slice2](https://blog.golang.org/go-slices-usage-and-internals_slice-2.png)

**grow slice**

```go
s1 := make([]byte, 5)
s2 := s1[2:4]
// grow s2 to its capacity by slicing it again.
s2 := s2[:cap(s2)]
```

A slice cannot be grown beyond its capacity.

To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.

```go
s := make([]int, 5)
t := make([]int, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 1
copy(t, s)
s = t
```

A common operation is to append data to the endo of a slice.

```go
func AppendByte(s []byte, data ...byte) []byte {
  m := len(s)
  n := m + len(data)
  if n > cap(s) {
    newSlice := make([]byte, (n+1)*2)
    copy(newSlice, s)
    s = newSlice
  }
  s = s[0:n]
  copy(s[m:n], data)
  return s
}
```

append func: `func append(s []T, x ...T) []T`.

usage:

```go
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) 
// equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

