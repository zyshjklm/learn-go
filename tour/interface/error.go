package main 

import (
	"fmt"
	"time"
	"math"
)


type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// --- Sqrt ---

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g\n", float64(e))

}

// the second return value should be error type.
// to call the Error func of ErrNegativeSqrt, return ErrNegativeSqrt(f)
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -0.0, ErrNegativeSqrt(f)
	}
	r := 1.0
	for {
		newton := r - (r*r - f) / (2*r)
		delta := math.Abs(newton - r)
		// fmt.Printf("tmp delta: %g, newton: %g\n", delta, newton)
		if delta > 1e-10 {
			r = newton
		} else {
			return r, nil
		}
	}
}



func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// sqrt
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
