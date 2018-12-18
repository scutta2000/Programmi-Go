package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []float64{-2, 0, 4, -9, 5, 6, 11, 8}
	min, max := minmax(slice)
	fmt.Printf("min: %f, max: %f", min, max)
}
func minmax(s []float64) (min, max float64) {
	min, max = s[0], s[0]
	if len(s) > 1 {
		mi, ma := minmax(s[1:])
		min = math.Min(min, mi)
		max = math.Max(max, ma)
	}
	return
}
