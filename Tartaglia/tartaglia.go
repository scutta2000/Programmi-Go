package main

import (
	"fmt"
	"os"
	"strconv"
)

//definire la funzione Tartaglia che ritorna il triangolo di Tartaglia di dimensione n
//la riga i-esima (i nel range [0,n-1]) contiene i coefficienti dello sviluppo di (a+b)^i
//si veda anche https://it.wikipedia.org/wiki/Triangolo_di_Tartaglia

func Tartaglia(n uint) [][]uint{
	tart:=make([][]uint,n)
	for i:=0;i<int(n);i++{
		tart[i]=make([]uint,i+1)
		tart[i][len(tart[i])-1]=1
		for y:=0;y<i+1;y++{
			if y==0||y==i{
				tart[i][y]=1
			}else{
				tart[i][y]=tart[i-1][y]+tart[i-1][y-1]
			}
		}
	}
	return tart
}

//visualizza il triangolo di Tartaglia di dimensione n, con n passato da linea di comando
func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: tartaglia <n>")
		return
	}
	if n, err :-= strconv.Atoi(os.Args[1]); err == nil {
		PrintSlice2D(Tartaglia(uint(n)))
	} else {
		fmt.Println("not a number")
	}
}

func PrintSlice2D(slice2 [][]uint) {
	for i := range slice2 {
		for y := range slice2[i] {
			fmt.Printf("%4d", slice2[i][y])
		}
		fmt.Println()
	}
}
