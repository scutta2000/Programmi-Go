package main

import (
	"fmt"
)

func clessidra(n, t int, flag bool) {
	if n == 0 {
		flag = true
	}
	c := 1
	if flag {
		c = -1
		if t == 0 {
			return
		}
	}
	draw(n, t)
	clessidra(n, t, flag)
	n -= c
	t += c
}

/// TODO: non funizona :(
func draw(n, t int) {
	for i := 0; i < t; i++ {
		fmt.Print(" ")
	}
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}
func main() {
	clessidra(3, 0, false)
}
