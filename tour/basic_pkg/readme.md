## basic

### 1. packages

packages

by convention, the package name is the same as the last element of the import path. "math/rand"

imports

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


variables

the var statement declares a list of variables, the type is last.
a var statement can be at package or function level.

a var declaration can include initializers. if an initializer is present, the type can be omitted; the variable will take the type of the initializer.

**inside a function**, the := short assignment statement can be used in place of a `var` declaration with implicit type.

