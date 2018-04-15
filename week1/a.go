package main

import (
	"fmt"
	"math"
)

func main() {
	Pow := math.Pow
	var c1x, c1y, c1r int
	var c2x, c2y, c2r int

	fmt.Scanf("%d %d %d", &c1x, &c1y, &c1r)
	fmt.Scanf("%d %d %d", &c2x, &c2y, &c2r)

	d := math.Sqrt(Pow(float64(c2x-c1x), 2) + Pow(float64(c2y-c1y), 2))

	if d < math.Abs(float64(c2r-c1r)) {
		fmt.Println(0)
	} else if d == math.Abs(float64(c2r-c1r)) {
		fmt.Println(1)
	} else if math.Abs(float64(c2r-c1r)) < d && d < math.Abs(float64(c1r+c2r)) {
		fmt.Println(2)
	} else if d == math.Abs(float64(c1r+c2r)) {
		fmt.Println(3)
	} else if d > math.Abs(float64(c1r+c2r)) {
		fmt.Println(4)
	}
}
