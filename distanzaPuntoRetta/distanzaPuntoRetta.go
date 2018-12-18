package main

import "fmt"
import "math"

func main() {
	var ax, ay, bx, by, m1, q1, m2, q2, d, dist float64
	fmt.Print("Inserisci le coordinate del punto poi l'm della retta poi il q della retta in fine la distanza ")
	fmt.Scan(&ax, &ay, &m1, &q1, &d)

	m2 = -1 / m1
	q2 = -m2*ax + ay
	bx = (q2 - q1) / (m1 - m2)
	by = bx*m2 + q2
	dist = math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	if dist < d {
		fmt.Println("il punto è a distanza minore di ", d)
	} else {
		fmt.Println("il punto è a distanza maggiore di ", d)
	}
}
