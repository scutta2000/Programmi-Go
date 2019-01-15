package main

// import "fmt"

//tipi usati per trovare le occorrenze di parole

type (
	occurrences map[string]int //mappa delle occorrenze di parole

	// da usare nella seconda parte dell'esercizio
	/*occurrence struct { // tipo occorrenza --da usare come tipo di una slice
		word string
		n    int
	}*/
)

//fare in modo che i tipi occurrences e occurrence implementino l'interfaccia fmt.Stringer
