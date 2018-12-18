package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Quanto alto deve essere il triangolo ")
	fmt.Scan(&n)
	drawTriangle(n)
}
func drawTriangle(n int) {
	var flag bool
	if n%2 == 1 {
		flag = true
	}
	n /= 2
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	if flag {
		for j := 0; j < n+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
