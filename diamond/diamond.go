package main

import (
	"fmt"
	"strings"
)

func drawDiamond(letter rune) {
	fmt.Println(letter)
	n := int(letter-'a')*2 + 1
	fmt.Println(n)
	for r := 0; r < n; r++ {
		if r <= n/2 {
			for i := 0; i < n/2-r; i++ {
				fmt.Print(" ")
			}
			fmt.Printf("%c", 'A'+r)
			if r != 0 {
				for i := 0; i < 2*r-1; i++ {
					fmt.Print(" ")
				}
				fmt.Printf("%c", 'A'+r)
			}
			fmt.Println()
		} else {
			for i := 0; i < r-n/2; i++ {
				fmt.Print(" ")
			}
			fmt.Printf("%c", n-r-1+'A')
			if r != n-1 {
				for i := 0; i < n/2-(r-n/2); i++ {
					fmt.Print(" ")
				}
				fmt.Printf("%c", n-r-1+'A')
			}
			fmt.Println()
		}
	}
}

func main() {
	var letter string
	fmt.Print("inserisci una lettera: ")
	fmt.Scan(&letter)
	fmt.Printf("%c", letter)
	drawDiamond(rune(strings.ToLower(letter)[0]))
}
