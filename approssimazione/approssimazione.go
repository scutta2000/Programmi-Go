package main

import "fmt"

func main() {
	var daApprossimare float64
	var approssimato int
	fmt.Print("Inserisci il numero che vuoi approssimare ")
	fmt.Scan(&daApprossimare)
	approssimato = int(daApprossimare)
	if (daApprossimare - float64(approssimato)) >= 0.5 {
		approssimato++
	}
	fmt.Println("Il numero è stato approssimato a", approssimato)
}
