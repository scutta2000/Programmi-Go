package main

import "fmt"

func main() {
	var aFloat float64
	var aInt int
	fmt.Print("Inserisci il valore da controllare ")
	fmt.Scan(&aFloat)
	aInt = int(aFloat)
	if aFloat-float64(aInt) == 0 {
		fmt.Println("Il numero è anche un intero")
	} else {
		fmt.Print("il numero non è un intero ")
		if aFloat-float64(aInt) < 0.5 {
			fmt.Println("e si arrotonda a", aInt)
		} else {
			fmt.Println("e si arrotonda a ", aInt+1)
		}
	}
}
