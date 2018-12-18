// Stampa le parti di un path, ad es. se il path e  "/home/lorenzo/lab" stampa
//home
//lorenzo
//lab
// Se il path non contiene alcuna occorrenza di '/' viene stampato così com'è.
package main

import (
	"fmt"
	"strings"
)

const sep = "/"

func main() {
	var path string
	fmt.Print("digita una string? ")
	fmt.Scanln(&path)
	var from, to int
	for {
		to = strings.Index(path[from:], sep)
		if to<0{
			break
		}
		fmt.Println(path[from:to+from])
		from = from+to+1
	}
	fmt.Println(path[from:])
}
