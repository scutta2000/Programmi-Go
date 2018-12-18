package main

import "fmt"

func main() {
	var oraAttuale, minutiAttuali, oraAppuntamento, minutiRimasti, oraRimaste int
	fmt.Scan(&oraAttuale, &minutiAttuali)
	oraAppuntamento = 24
	minutiRimasti = (oraAppuntamento-oraAttuale)*60 - minutiAttuali
	oraRimaste = minutiRimasti / 60
	minutiRimasti = minutiRimasti % 60
	fmt.Print(oraRimaste, ":", minutiRimasti)
}
