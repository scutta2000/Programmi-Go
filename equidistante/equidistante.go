package main

import "fmt"
import "math"

func main() {
	var ax, ay, bx, by, ex, ey, distA, distB float64
	fmt.Print("inserisci le coorinate del primo punto ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Inserisci la lunghezza del secondo punto ")
	fmt.Scan(&bx, &by)
	fmt.Print("Inserisci le coordinate del punto da controllare ")
	fmt.Scan(&ex, &ey)
	distA = math.Sqrt(math.Pow(ax-ex, 2) + math.Pow(ay-ey, 2))
	distB = math.Sqrt(math.Pow(bx-ex, 2) + math.Pow(by-ey, 2))
	if math.Abs(distA-distB) <= 0.001 {
		fmt.Println("Il punto è equidistante ")
	} else {
		fmt.Println("Il punto non è equidistante")
	}

}
