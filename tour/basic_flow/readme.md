## basic 



### 2. flow

**for**

Go has only one looping construct. the `for` loop.

the basic `for` loop looks as it does in C or Java, except that the `( )` are gone (they are not even optional) and the `{ }` are required.

you can leave the pre and post statements empty. especially, drop the semicolons: C's `while` is spelled `for` in Go.

forever: if you omit the loop condition it loop forever. 

```go
for {
  // forever loop here! do sth in each loop!
}
```



**if**

the `if` statement looks as it does in C or Java, except that the `( )` are gone and the `{ }` are required.

if with a short statement: if statement can start with a short statement execute before the condition.

variables declared by the statement are only in scope untile the end of the `if`. so variables are also available inside any of the `else` blocks.



**exercise: loops and functions for square root**

```go
root := 1.0

root = root - (root*root - x) / (2*root)
```

