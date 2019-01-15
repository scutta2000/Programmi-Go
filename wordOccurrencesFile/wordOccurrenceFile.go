package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// visualizza le occorrenze di tutte le "parole" contenute nel file di testo specificato da linea di comando
// usando le funzioni definite di seguito
func main() {
	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Println("ERRORE", err)
		return
	}
	mappa, err := readOccurrencesFile(f)
	// for k, v := range mappa {
	// 	fmt.Printf("%-20s%d\n", k, v)
	// }
	printSortedByValue(mappa)
}

// dato un descrittore di file, crea la map occm delle occorrenze delle parole che compaiono nel file
// in caso di errore durante la lettura del file restituisce (nil,err), dove err è l'errore
// altrimenti ritorna (occm,nil)
// suggerimento: usare bufio.NewReader oppure bufio.NewScanner
func readOccurrencesFile(f *os.File) (occurrences, error) {
	var line string
	mappa := make(occurrences)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		out := splitLine(line)
		// fmt.Println(out)
		for _, c := range out {
			if c == "" {
				continue
			}
			mappa[strings.ToLower(c)]++
		}
	}

	return mappa, nil

}

// estrae le parole contenute in una stringa intendendo per parola una qualsiasi sequenza di caratteri latini
// comprese lettere accentate etc. (questi caratteri particolari hanno code points compresi tra i valori esadecimali c0 e ff)
// suggerimento: usare il package regexp (in particolare, il metodo MustCompile e ...)
func splitLine(s string) (words []string) {
	words = regexp.MustCompile(`[a-zA-Z\xc0-\xff]+`).FindAllString(s, -1)
	return
}

// seconda parte dell'esercizio: stampare la mappa delle occorrenze
// 1) in ordine alfabetico
// 2) in ordine rispetto al numero di occorrenze

// stampa una mappa di occorrenze di parole, in ordine alfabetico
// usare il package sort (in particolare, sort.Strings)
func printSortedByKey(occ occurrences) {
	var output []string
	for k, v := range occ {
		output = append(output, fmt.Sprintf("%-20s%d\n", k, v))
	}
	sort.Strings(output)
	fmt.Println(output)
}

// stampa una mappa di occorrenze di parole, in ordine di numero di occorrenze
// usare il package sort (in particolare, sort.Slice o sort.SliceStable)
func printSortedByValue(occ occurrences) {
	type occor struct {
		key   string
		value int
	}
	var output []occor
	for k, v := range occ {
		output = append(output, occor{k, v})
	}
	sort.SliceStable(output, func(i int, j int) bool { return output[i].value < output[j].value })
	for _, c := range output {
		fmt.Printf("%-15s%d\n", c.key, c.value)
	}
}
