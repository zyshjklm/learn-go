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


