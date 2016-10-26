## commentary

Go provides C-style /* */ block comments and C++-style // line comments.

Line comments are the norm; block comments appear mostly as package comments,
but are useful to disable large swaths of code.

godoc processes Go source to extract documentation about the contents 
of the package. comments that appear before top-level declarations, with no
intervening newlines, are extracted along with the declaration to serve as
explanatory text for the item.

Every package should have a package comment, a block comment preceding the
package clause(分句,条款). The package comment should introduce the package and provide infomation relavant the package as a whole.

Comments do not need extra formatting such as banners of stars.

godoc -- lick gofmt, take care of that.

the comments are uninterpreted plain text.

Inside a package, any comment immediately preceding(先于) a top-level 
declaration serve as a **doc comment** for that declaration. Every 
exported(Capitalized) name in a program should have a doc comment.

example

src/strconv/atob.go

```go
// FormatBool returns "true" or "false" according to the value of b
func FormatBool(b bool) string {
        if b {
                return "true"
        }
        return "false"
}

​```shell

godoc strconv | grep FormatBool
#     FormatBool, FormatFloat, FormatInt, and FormatUint convert values to
#     s := strconv.FormatBool(true)
# func FormatBool(b bool) string
#     FormatBool returns "true" or "false" according to the value of b

```

the syntax of declaration allows grouping of declarations.


### my works of compare docs and source code

* source code: github.com/go/src/strconv/
* docs: http://localhost:6060/pkg/github.com/go/src/strconv/

in docs, there are Overview, Index and Examples.

1. Overview, comes from doc.go file. 

2. Index extract from all source code files.

**const**

source code: 

```go
const intSize = 32 << (^uint(0) >> 63)

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize
```

comments:

```go
const IntSize = intSize
IntSize is the size in bits of an int or uint value.
```

**Variables**

variables like const.

**Function**

source code:

```go
// AppendBool appends "true" or "false", according to the value of b,
// to dst and returns the extended buffer.
func AppendBool(dst []byte, b bool) []byte {
        if b {
                return append(dst, "true"...)
        }
        return append(dst, "false"...)
}
```

comments:

```go
func AppendBool(dst []byte, b bool) []byte

AppendBool appends "true" or "false", according to the value of b, to dst and returns the extended buffer.
```

**example**

source code :  from **example_test.go**

```go
func ExampleAppendBool() {
        b := []byte("bool:")
        b = strconv.AppendBool(b, true)
        fmt.Println(string(b))

        // Output:
        // bool:true
}
```

comments example:

Code:

```
b := []byte("bool:")
b = strconv.AppendBool(b, true)
fmt.Println(string(b))
```

Output:

```
bool:true
```

