package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	zn := 1.0
	delta := 0.0000000001
	deltaX := (zn*zn - x) / (2 * zn)
	for math.Abs(deltaX) > delta {
		zn = zn - deltaX
		deltaX = (zn*zn - x) / (2 * zn)
	}
	return float64(zn), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
