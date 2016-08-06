// package tempconv performs Celsius and Fahrenheit temprature computations.
package tempconv

// import nothing

// convert a Celsius temprature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit( c*9/5 + 32)
}

// convert a Fahrenheit temprature to Celsius.
func FToC(f Fahrenheit) Celsius {
    return Celsius((f-32) * 5 / 9)
}

// convert a Celsius temprature to Kelvin.
func CToK(c Celsius) Kelvin {
    return Kelvin( c + (-AbsoluteZeroC))
}

// convert a Kelvin temprature to Celsius.
func KToC(k Kelvin) Celsius {
    return Celsius(k) + AbsoluteZeroC
}
