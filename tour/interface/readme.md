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



