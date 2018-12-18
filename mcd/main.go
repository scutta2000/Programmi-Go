package main

import (
	"fmt"
)

func mcdRic(m, n int) int {
	if m == 0 {
		return n
	}
	return mcdRic(n%m, m)
}
func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(mcdRic(m, n))
}
