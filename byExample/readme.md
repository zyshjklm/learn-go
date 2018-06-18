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

