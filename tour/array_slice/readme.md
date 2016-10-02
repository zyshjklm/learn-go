### 1 array

the type `[n]T` is an array of n values of type `T`.

```go
var a [10]int
# declares a variable `a` as an array of ten integers.
```

an array's length is part of its type. so array cannot be resized.



### 2 Slice

a slice points to an array of values and also includes a length.

`[]T` is a slice with element of type `T`.

```go
p := []int{2, 3, 5, 7, 11, 13}
fmt.Println("p ==", p)
fmt.Println("p[1:4] ==", p[1:4])
```

slice can be re-sliced, creating a new value that points to the same array.

`s[low:high]` evaluates to a slice of the elements from `low` through `high-1`.

`s[lo:lo]` has no element, and `s[lo:lo+1]` has one element.

**makeing slices**

slices are created with the `make` function. it works by allocating a zeroed array and returning a slice that refers to that array.

to specify a capacity, pass a third argument to `make`.

```go
a := make([]int, 5) 	// len(a) = 5
b := make([]int, 5, 8)	// len(a) = 5, cap(b) = 8
```

**nil slices**

the zero value of a slice is `nil`. a nil slice has a length and capacity of 0.

**adding elements to a slice**

built-in `append` function.

```go
func append(s []T, vs ...T) []T
// s is a slice of type T. 
// the rest are T values to append to the slice
```

if the backing array of s is too small to fit all the given values a bigger array will be allocated.



### 3 go slice usage and internals

refer: https://blog.golang.org/go-slices-usage-and-internals

#### 3.1 array

[4]int. represents an array of four integers.

an array's size is fixed.

arrays do not need to be initialized explicitly. **the zero value of an array is a read-to-use** whose elements are themselves zeroed.

the in-memory representation of [4]int is just for integer values laid out sequentially.

Go's arrays are values. an array variable denotes the entire array. it is not a pointer to the first element. this means that when you assign or pass around an array value, you will make a copy of its contents.

to avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.

one way to think about arrays is as a sort of struct but with indexed rather than named fields.

#### 3.2 slice

array have their place, but they're a bit inflexible. 

slice, through, are everywhere. the build on arrays to provide great power and convenience.

`[]T` slice has no spcified length.

a slice can be created with the built-in function called make.

```go
func make( []T, len, cap )  []T
// cap is optional capacity. 
```

slice is done by specifying a half-open range with two indices separated by a colon.

#### 3.3 slice internals

a slice is a **descriptor** of an array segment. it consists of a **pointer** to the array, the **length** of the segment, and its **capacity** (the max length of the segment).

slice:

* pointer. *Elem
* len. int. the number of elements referred to by the slice.
* cap. int. the number of elements in the underlying array.

make a slice.

```go
make( []byte, 5)
```

**slicing** does **not copy** the slice's data. it creates a new slice value that points to the original array.

a slice cannot be grown beyond its capacity. attempting to do so will cause a runtime panic.

slices cannot be re-sliced below zero to access earlier elements in the array.

**growing slices**

扩大切片的容量，需要新建一个更大的切片，并将原来的数据拷贝到新的切片中。这就是其它语言中，动态数组的实现技术。

doubles the capacity of s :

```go
t := make( []byte, len(s), (cap(s)+1) *2 ) 
// +1 in case cap(s) == 0
for i := range s {
  t[i] = s[i] 
}
s = t
```

append operation.

```go
// append bytes
func AppendByte(slice []byte, data ...byte) []byte {
  m := len(slice)
  n := m + len(data)
  if n > cap(slice) { 
    // allocate double what's needed, for future growth.
    newSlice := make( []byte, (n+1)*2)
    copy(newSlice, slice)
    slice = newSlice
  }
  slice = slice[0:n]	// slice length
  copy(slice[m:n], data)
  return slice
}
```

built-in `append` function :

```go
func append( s []T, x ...T) []T
```









