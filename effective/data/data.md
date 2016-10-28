## Data



### 1 Allocation with new

Go has two allocation primitives, the built-in function `new` and `make`.

`new` allocate memory but does not `initialize` the memory, it only `zeros` it.

new(T): 

* allocate zeroed storage for a new item of type T, and
* return its address, a value of type *T, 

*T is a pointer to a newly allocated zero value of type T.

Since  the memory returned by new is zeroed, a user of the data structure can create one with new and get right to work, without further initialization.

zero value:

* bytes.Buffer: an empty buffer ready to use
* sync.Mutex: unlocked mutex

example:

```go
type SyncedBuffer struct {
  	lock 	sync.Mutex
  	buffer	bytes.buffer
}

p := new(SyncedBuffer)	// type *SyncedBuffer
var v SyncedBuffer		// type SyncedBuffer
```



### 2 Constructors and composite literals

sometimes the zero value is not good enough and an initializing constructor is necessary.

derived example:

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

There is a lot of boiler-plate(公式化的，陈腐的) in there. We can simplify it using a `composite literal`, which is an expression that create a new instance each it is evaluated.

```go
func NewFile(fd int, name string) *File {
  	if fd < 0 {
      	return nil
  	}
  f := File{fd, name, nil, 0}
  return &f
}
```

It is perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns.

The last two line can combined as below:

```go
return &File{fd, name, nil, 0}
```

The fields of a composite literal are laid out in order and muste all be present. By labeling the elements explicitly as fileld: value pairs, the initializers can appear in any order, with the missing ones left as their respective zero values.

```go
return &File{fd: fd, name: name}
```

A limiting case. Type{} contains no fields at all, it create a zero value for the type.

so new(File) and &File{} are equivalent.

Compsite literals can also be create for arrays, slices and maps.

```go
// array
a := [...]string   {"no error", "Eio", "invalid argument"}
// slice
s := []string   {"no error", "Eio", "invalid argument"}
m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}
```





















