package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	zn := 1.0
	delta := 0.0000000001
	deltaX := (zn*zn - x) / (2 * zn)
	for math.Abs(deltaX) > delta {
		zn = zn - deltaX
		deltaX = (zn*zn - x) / (2 * zn)
	}
	return float64(zn)
}

func main() {
	fmt.Println("programmd sqrt result:", Sqrt(2))
	fmt.Println("built-in function result:", math.Sqrt(2))
}
