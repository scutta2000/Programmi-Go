package main

import "fmt"

func main() {
	var ax, ay, bx, by, px, py int
	fmt.Print("Inserisci le coordinate del primo vertice ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Inserisci le coordinate del secondo vertice ")
	fmt.Scan(&bx, &by)
	fmt.Print("Inserisci le coordinate del punto da analizzare ")
	fmt.Scan(&px, &py)
	if px > ax && px < ax && py < ay && py > by {
		fmt.Println("P è dentro il rettangolo")
	} else {
		fmt.Println("P è fuori dal rettangolo")
	}
}
