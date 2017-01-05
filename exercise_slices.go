package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pix := make([][]uint8, dy)
	for i := 0; i < dy;  i++ {
		pix[i] = make([]uint8, dx)
	}
	
	for i, v := range pix {
		for j, subV := range v {
			subV = uint8((i + j) / 2)
			pix[i][j] = subV
		}
	}
	return pix
}

func main() {
	pic.Show(Pic)
}
