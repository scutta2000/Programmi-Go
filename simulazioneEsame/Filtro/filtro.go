package main

import "fmt"

func main() {
	var input, tra18_21, tra22_24, tra25_27, tra28_30 float64
	fmt.Scan(&input)
	for input != 0 {
		switch {
		case input < 18:
			break
		case input > 30:
			break
		case input <= 21:
			tra18_21++
		case input <= 24:
			tra22_24++
		case input <= 27:
			tra25_27++
		case input <= 30:
			tra28_30++
		}
		fmt.Scan(&input)
	}
	tot := tra18_21 + tra22_24 + tra25_27 + tra28_30
	fmt.Printf("A : %.2f %%\n", tra28_30/tot*100)
	fmt.Printf("B : %.2f %%\n", tra25_27/tot*100)
	fmt.Printf("C : %.2f %%\n", tra22_24/tot*100)
	fmt.Printf("D : %.2f %%\n", tra18_21/tot*100)
}
