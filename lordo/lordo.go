package main

import "fmt"

func main() {
	var lordo, iva, imponibile float64
	fmt.Scan(&lordo)
	fmt.Scan(&iva)
	imponibile = (lordo * 100.0) / (iva + 100.0)
	fmt.Print(imponibile)
}
