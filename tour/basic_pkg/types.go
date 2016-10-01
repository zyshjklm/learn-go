package main 

import (
	"fmt"
	"math/cmplx"
	"math"
)

const PI = 3.14159

// Numeric constants are high-precision values.
const (
	Big		= 1 << 100
	Small	= Big >> 99
)

var (
	ToBe	bool	= false
	MaxInt uint64	= 1<<64 - 1
	z 	complex128	= cmplx.Sqrt(-5 + 12i)
)

func zeroValue() {
	// zero value
	var i int 
	var f float64
	var b bool
	var s string
	fmt.Printf("zero values: %v,%v,%v,%v\n", i, f, s, b)
}

func sqrt() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x*0.1 }

func main() {
	// %T type;  %v value.
	const fomt = "%T(%v)\n"

	fmt.Printf(fomt, ToBe, ToBe)
	fmt.Printf(fomt, MaxInt, MaxInt)
	fmt.Printf(fomt, z, z)

	zeroValue()
	sqrt()

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", PI, "Day")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	
	// constant 1267650600228229401496703205376 overflows int
	//fmt.Println(needInt(Big))
}

