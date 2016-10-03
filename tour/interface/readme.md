## interface

an interface type is defined by a set of methods.

a value of interface type can hold any value that implements those methods.

**interfaces are satisfied implicitly.**

a type implements an interface by implementing the methods. there is no explicit declaration of intent.

implicit interfacers decouple implementation packages from the packages that define the interfaces: neither depends on the other.



**Stringers**

one of the most ubiquitous interface is `Stringer` defined by the `fmt` package.

```go
type Stringer interface {
   String() string
}
```

a `Stringer ` is a type that can describe itself as a string. the `fmt` package (and many others ) look for this interface to print values.

 

IPAddr Stringers.

```go
IPAddr{8, 8, 4, 4}
```

byte to string:

* byte to int
* use striconv.Itoa() to trans int to string.



#### Errors

Go programs express error state with `error` values.

the `error` type is a built-in interface similar to `fmt.Stringer`:

```go
type error interface {
   Error() string
}
```

function often return an `error` value, and callling code should handle errors by testing whether the error equals `nil`.

```go
i, err := strconv.Atoi("42")
if err != nil {
  fmt.Printf("couldn't convert number: %v\n", err)
}
fmt.Println("Converted integer:", i)
```

a nil error denotes success.