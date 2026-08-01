package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		old := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(old-z) < 1e-10 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
