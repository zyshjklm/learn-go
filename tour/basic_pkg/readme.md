## basic

### 1. packages

**packages**

by convention, the package name is the same as the last element of the import path. "math/rand"

**imports**

```go
// imported into a parenthesized statement.
import (
    "fmt"
    "math"
)

// like
import "fmt"
import "math"
```

exported name.

after importing a package, you can refer to the name it exports.
a name is exported if it begins with a capital letter, like math.Pi.


functions

* zero or more arguments. 
* the type comes after the variable name.
* two or more consecutive named func parameters share a type, you can omit the type from all but the last.
* can return any number of results.
* return values may be named and act just like variables. -- for short function.


**variables**

the var statement declares a list of variables, the type is last.
a var statement can be at package or function level.

a var declaration can include initializers. if an initializer is present, the type can be omitted; the variable will take the type of the initializer.

**inside a function**, the := short assignment statement can be used in place of a `var` declaration with implicit type.



**basic types**

basic type of Go:

```go
bool
string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point

float32 float64
complex64 complex128
```

rune，如尼文，是一套字母表。

variables declared without an explicit initial value are given their zero value.

* `0` for numeric types

* `false` for boolean type

* `""` (empty string) for strings

type conversions:

T(v): conversion the value `v` to the type `T`.

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// or short 
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion.



**Constans**

constants are declared like variables, but with the `const` keyword.

constants can be character, string, bool or numeric values.

constant cannot be decalred use the `:=` syntax.



Numeric constants are high-precision *values*. see `Big` const of variables.go.





   