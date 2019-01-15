package main

import (
	"fmt"
)

func main() {
	p := make(Nome2Voti)
	p["Paolo Boldi"] = []Voti(18, 24, 28, 19)
	p["Giovanni Rossi"] = []Voti(30)
	fmt.Print(p)
}
