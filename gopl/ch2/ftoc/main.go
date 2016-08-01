// Ftoc print Fahrenheit-to-Celsius conversions
package main 

import (
	"fmt"
)

func main() {
	const FreezingF, boilingF = 32.0, 212.0
	// 32°F = 0°C
	fmt.Printf("%g°F = %g°C\n", FreezingF, FToC(FreezingF))
	// 212°F = 100°C
	fmt.Printf("%g°F = %g°C\n", boilingF, FToC(boilingF))	
}

func FToC(f float64) float64 {
	return (f - 32 ) * 5 / 9
}