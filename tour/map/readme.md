### map

a map maps keys to values.

maps muste be created with make ( not new) before use. the `nil` map is empty an cannot be assigned to.

mutating maps:

```go
// insert or update
m[key] = elem

// retrieve 
elem = m[key]

// delete
delete(m, key)

// test that a key is present with a two-vaue assignment
elem, ok = m[key]
```

if `key` is in `m`, `ok` is `true`, if not, `ok` is `false` and `elem` is the zero value for the map's element type.

