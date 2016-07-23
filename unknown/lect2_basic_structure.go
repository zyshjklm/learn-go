// 注意下面各部分的顺序

package main

// import "fmt"
import stdio "fmt"
//import . "fmt"  // fmt.Println -> Println

const PI = 3.14

// global var
var name = "gohper"

type newType int

// struct
type gopher struct{}

// interface
type golang interface{}

func main() {
//    fmt.Println("Hello world!")
    stdio.Println("Hello world!")
}
// 可见性：首字母的大小写，区分可见性。大写对包外是可见的
