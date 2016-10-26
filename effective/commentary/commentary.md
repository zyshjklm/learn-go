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

```shell

godoc strconv | grep FormatBool
#     FormatBool, FormatFloat, FormatInt, and FormatUint convert values to
#     s := strconv.FormatBool(true)
# func FormatBool(b bool) string
#     FormatBool returns "true" or "false" according to the value of b

```

the syntax of declaration allows grouping of declarations.

