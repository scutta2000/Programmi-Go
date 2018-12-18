/*
input x
controllo che sia compreso tra 1 e 99
    seno viene terminato il programma
stampo in lettere il numero
*/

package main

import "fmt"

func main() {
	var x int
	fmt.Print("inserisci un numero compreso tra 1 e 99 ")
	fmt.Scan(&x)
	t := ""
	if x > 99 || x < 0 {
		fmt.Print("Sei stupido ho detto compreso tra 1 e 99")
	} else if x >= 10 && x < 20 {
		switch x {
		case 10:
			t = "dieci"
		case 11:
			t = "undici"
		case 12:
			t = "dodici"
		case 13:
			t = "tredici"
		case 14:
			t = "quattordici"
		case 15:
			t = "quindici"
		case 16:
			t = "sedici"
		case 17:
			t = "diciasette"
		case 18:
			t = "diciannove"
		}
	} else {
		v := "a"
		switch x / 10 {
		case 2:
			t = "vent"
			v = "i"
		case 3:
			t = "trent"
		case 4:
			t = "quarant"
		case 5:
			t = "cinquant"
		case 6:
			t = "sessant"
		case 7:
			t = "settant"
		case 8:
			t = "ottant"
		case 9:
			t = "novant"
		}
        var t1 string
		switch x % 10 {
		case 1:
			t1 = "uno"
            v=""
		case 2:
			t1 = "due"
		case 3:
			t1 = "tre"
		case 4:
			t1 = "quattro"
		case 5:
			t1 = "cinque"
		case 6:
			t1 = "sei"
		case 7:
			t1 = "sette"
		case 8:
			t1 = "otto"
            v=""
		case 9:
			t1 = "nove"
		}
        t += v
        t+=t1
	}
	fmt.Printf("%s",t)
}
