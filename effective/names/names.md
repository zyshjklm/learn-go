## names

Names in Go have sematic effect: the visibility of a name outside a package in determined by whether its first character is upper case.

### 1 package names

when a package is imported, the package name becomes an accessor for the 
contents. It is helpful if everyone using the package can use the same name 
to refer to its contents, which implies that the package name should be good:
short, concise, evocative(易唤起记忆的).

by convention, packages are given lower case, single-word names.

a nother convention is that the package name is the base name of its source directory.

the package in src/encoding/base64 is imported as "encoding/base64" but has name base64.

The importer of a package will use the name to refer to its contents.

the buffered reader in the bufio package is called Reader, the caller see it as bufio.Reader.

tips:

a helpful doc comment can often be more valuable than an extra long name.

### 2 Getters

Go does not provide automatic support for getters and setters.

If you nave a field called `owner` , the getter method should be called `Owner()`, not

`GetOwner()`.  The use of upper-case name for export provides the hook to 

discriminate(区别对待) the field from the method.

A setter function, if needed, will likely be called `SetOwner`.

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### 3 Interface names

By convention, one-method interface are named by the method plus an 
-er suffix or similar modification to construct an agent noun: Reader, 
Formatter, 

method: Read
interface: Reader

### 4 MixedCaps
in Go, use MixedCaps or mixedCaps rather than underscores to write multiword names.
