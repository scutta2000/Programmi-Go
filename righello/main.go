package main

import "fmt"

func righello(n int) {
	if n == 0 {
		return
	}
	righello(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	righello(n - 1)
}
func main() {
	righello(6)
}
