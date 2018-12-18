package main

import "fmt"

func main() {
	var m, q, x, y float64
	fmt.Print("Inserisci il i parametri m e q della retta")
	fmt.Scan(&m, &q)
	fmt.Print("Inserisci le coorinate del punto")
	fmt.Scan(&x, &y)
	if y-m*x+q <= 0.0001 {
		fmt.Println("Il punto appartiene alla retta")

	} else {
		fmt.Println("Il punto non appartiene alla retta")
	}
}
