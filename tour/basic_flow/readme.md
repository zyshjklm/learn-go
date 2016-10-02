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



**switch**

a case body breaks automatically, unless it ends with a `fallthrough` statement.

switch cases evaluate from top to bottom, stopping when a cases succeeds.

```go
switch i {
  case 0:
  	fmt.Println("it 0")
  case f():
  	// does not call f() if i == 0.
    fmt.Println("mit f()")
}
```

switch without a condition is the same to `switch true`.



**defer**

a `defer` statement defers the executiong of a function until the surrounding function returns.

the deferred call's **arguments** are evaluated immediately. but the function call is not executed until the surrounding function returns.

deferred function calls are pushed onto a stack. when the function returns, it deferred calls are executed in last-in-first-out order.

### 2.1 Defer, Panic and Recover

refer: https://blog.golang.org/defer-panic-and-recover

#### defer

a **defer statement** pushes a function call onto a list. the list of saved calls is executed after the surrounding function returns. 

defer is commonly used to simplify functions that perform various clean-up actions.

for example, use defer to close file, defer statements allow us to think about closing each file right after opening it, guaranteeing that, regardless of the number of return statement in the function, the file `will` be closed.

the behavior of defer is straighforward and predictable(直接了当，且可预测）. there are three simple rules:

- a defer function's arguments are evaluated when the defer statement is evaluated.
- deferred function calls are executed in Last-in-First-out order after the surrounding function returns
- deferred functions may read and assign to the returning function's named return values.

rule 3 is convenient for modifying the error return code of a function.

#### Panic and Recover 

**panic** is a built-in function that stops the ordinary flow of control and begins `panicking`.

when the func F calls panic, execution of F stops, any deferred func in F are executed normally, and then F return to its caller.

the process continues up the stack until all functions in the current gorountine have returned, at which point the program crashes. 内层函数的panic会沿着stack一直向上层传导，层层传递panic异常信息，最终导致最外层的函数崩溃。

**recover** is a built-in function that regains control of a panicking gorouting. recover is only useful inside deferred functions.

if the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.



defer usage: 

* file
  * f = os.Open(xxx)
  * defer f.Close()
* mutex
  * mu.Lock()
  * defer mu.Unlock()
* header & footer
  * printHeader()
  * defer printFooter()