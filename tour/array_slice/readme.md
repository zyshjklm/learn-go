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

