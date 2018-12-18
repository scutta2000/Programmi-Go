package main

import "fmt"

func main() {
	var ora1, min1, ora, ora2, min2 int
	fmt.Println(" inserisci ora attuale")
	fmt.Scan(&ora1, &min1)
	ora = (23-ora1)*60 + (60 - min1)
	ora2 = ora / 60
	min2 = ora - (ora2 * 60)
	fmt.Println(ora2, min2)
}
