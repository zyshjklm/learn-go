### functions



functions are values too.

```go
hypot = func(x, y float64) float64 {
  return math.Sqrt(x*x + y*y)
}
```

Go functions may be closures. A closure is a function value that references variables from outside its body.

the function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

