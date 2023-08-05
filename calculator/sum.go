package calculator

import "fmt"

var offset float64 = 1 // private
var Offset float64 = 1 // public

func Sum(a float64, b float64) float64 { // public
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("multiply: ", Multiply(a, b))
	return a + b + offset
}
