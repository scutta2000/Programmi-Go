package main

import "fmt"
import "math"

func main() {
	var cx, cy, r, px, py, dist float64
	fmt.Print("inserisci le coorinate del centro ")
	fmt.Scan(&cx, &cy)
	fmt.Print("Inserisci la lunghezza del raggio ")
	fmt.Scan(&r)
	fmt.Print("Inserisci le coordinate del punto da controllare ")
	fmt.Scan(&px, &py)
	dist = math.Sqrt(math.Pow(cx-px, 2) + math.Pow(cy-py, 2))
	if dist < r {
		fmt.Print("Il punto è dentro il cerchio ")
	} else {
		fmt.Println("Il punto è fuori dal centro")
	}

}
