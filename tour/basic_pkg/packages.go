package main 

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("My favorite number is :", rand.Intn(50))
	}
	// the same result for each running.

	// update seed for each time.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		fmt.Println("My favorite number is", r.Intn(50))
	}

	// Nextafter() returns the next representable value after x towards y.
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))

	fmt.Println("PI is:", math.Pi)
}